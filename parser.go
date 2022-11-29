// Copyright 2022 eatmoreapple.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redis

import (
	"bufio"
	"context"
	"errors"
	"io"
	"strconv"
	"time"
)

type Parser interface {
	Parse(ctx context.Context, reader io.Reader) error
}

func NewProtocolParser(reader io.Reader) *ProtocolParser {
	return &ProtocolParser{rd: bufio.NewReader(reader)}
}

type ProtocolParser struct {
	rd *bufio.Reader
}

func (p ProtocolParser) Parse() ([]byte, error) {
	line, isPrefix, err := p.rd.ReadLine()
	if err != nil {
		return nil, err
	}
	if isPrefix {
		return nil, errors.New("line too long")
	}
	prefix := line[0]
	switch prefix {
	case StatusReply:
		return line[1:], nil
	case ErrorReply:
		return nil, errors.New(string(line[1:]))
	case IntReply:
		return line[1:], nil
	case StringReply:
		size, err := parseInt(line[1:])
		if err != nil {
			return nil, err
		}
		if size == -1 {
			return nil, ErrNil
		}
		ps := make([]byte, size+2)
		_, err = p.rd.Read(ps)
		if err != nil {
			return nil, err
		}
		return ps[:size], nil
	case ArrayReply:
		size, err := parseInt(line[1:])
		if err != nil {
			return nil, err
		}
		var buf []byte
		for i := 0; i < size; i++ {
			p, err := p.Parse()
			if err != nil {
				return nil, err
			}
			buf = append(buf, p...)
		}
		return buf, nil
	}
	return nil, errors.New("invalid reply type")
}

func parseInt(bytes []byte) (int, error) {
	if len(bytes) == 0 {
		return 0, errors.New("invalid number")
	}
	return strconv.Atoi(string(bytes))
}

// StringParser implement
type StringParser struct {
	v   string
	err error
}

func (s *StringParser) Parse(ctx context.Context, reader io.Reader) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		parser := NewProtocolParser(reader)
		v, err := parser.Parse()
		if err != nil {
			return err
		}
		s.v = string(v)
		return nil
	}
}

func (s *StringParser) SetErr(err error) {
	s.err = err
	if err != nil {
		s.v = ""
	}
}

func (s *StringParser) String() string {
	return s.v
}

func (s *StringParser) StringResult() *StringResult {
	return &StringResult{
		val: s.v,
		err: s.err,
	}
}

func (s *StringParser) BoolResult() *BoolResult {
	return &BoolResult{
		val: s.v == "OK",
		err: s.err,
	}
}

type IntegersParser struct {
	v   int64
	err error
}

func (i *IntegersParser) Parse(ctx context.Context, reader io.Reader) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		parser := NewProtocolParser(reader)
		v, err := parser.Parse()
		if err != nil {
			return err
		}
		count, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return err
		}
		i.v = count
		return nil
	}
}

func (i *IntegersParser) IntResult() *IntResult {
	return &IntResult{
		val: i.v,
		err: i.err,
	}
}

func (i *IntegersParser) DurationResult() *DurationResult {
	return &DurationResult{
		val: time.Duration(i.v) * time.Second,
		err: i.err,
	}
}

func (i *IntegersParser) SetErr(err error) {
	i.err = err
	if err != nil {
		i.v = 0
	}
}

func (i *IntegersParser) BoolResult() *BoolResult {
	return &BoolResult{
		val: i.v > 0,
		err: i.err,
	}
}

type StringSliceParser struct {
	v   []string
	err error
}

func (s *StringSliceParser) Parse(ctx context.Context, reader io.Reader) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		rd := NewProtocolParser(reader)
		line, isPrefix, err := rd.rd.ReadLine()
		if err != nil {
			return err
		}
		if isPrefix {
			return ErrLineTooLong
		}
		prefix := line[0]
		switch prefix {
		case ErrorReply:
			return errors.New(string(line[1:]))
		case ArrayReply:
			size, err := parseInt(line[1:])
			if err != nil {
				return err
			}
			s.v = make([]string, size)
			for i := 0; i < size; i++ {
				p, err := rd.Parse()
				if err != nil {
					return err
				}
				s.v[i] = string(p)
			}
			return nil
		default:
			return ErrInvalidReply
		}
	}
}

func (s *StringSliceParser) Strings() []string {
	return s.v
}

func (s *StringSliceParser) SetErr(err error) {
	s.err = err
	if err != nil {
		s.v = nil
	}
}

func (s *StringSliceParser) StringSliceResult() *StringSliceResult {
	return &StringSliceResult{
		val: s.v,
		err: s.err,
	}
}

type FloatsParser struct {
	v   float64
	err error
}

func (f *FloatsParser) Parse(ctx context.Context, reader io.Reader) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		parser := NewProtocolParser(reader)
		v, err := parser.Parse()
		if err != nil {
			return err
		}
		f.v, err = strconv.ParseFloat(string(v), 64)
		return err
	}
}

func (f *FloatsParser) FloatResult() *FloatResult {
	return &FloatResult{
		val: f.v,
		err: f.err,
	}
}

func (f *FloatsParser) SetErr(err error) {
	f.err = err
	if err != nil {
		f.v = 0
	}
}

type StringMapParser struct {
	v   map[string]string
	err error
}

func (s *StringMapParser) Parse(ctx context.Context, reader io.Reader) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		rd := NewProtocolParser(reader)
		line, isPrefix, err := rd.rd.ReadLine()
		if err != nil {
			return err
		}
		if isPrefix {
			return ErrLineTooLong
		}
		prefix := line[0]
		switch prefix {
		case ErrorReply:
			return errors.New(string(line[1:]))
		case ArrayReply:
			size, err := parseInt(line[1:])
			if err != nil {
				return err
			}
			s.v = make(map[string]string, size)
			for i := 0; i < size/2; i++ {
				key, err := rd.Parse()
				if err != nil {
					return err
				}
				value, err := rd.Parse()
				if err != nil {
					return err
				}
				s.v[string(key)] = string(value)
			}
			return nil
		default:
			return ErrInvalidReply
		}
	}
}

func (s *StringMapParser) StringMap() map[string]string {
	return s.v
}

func (s *StringMapParser) SetErr(err error) {
	s.err = err
	if err != nil {
		s.v = nil
	}
}

func (s *StringMapParser) StringMapResult() *StringMapResult {
	return &StringMapResult{
		val: s.v,
		err: s.err,
	}
}

type ScanParser struct {
	v      []string
	cursor uint64
	err    error
}

func (s *ScanParser) Parse(ctx context.Context, reader io.Reader) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		rd := NewProtocolParser(reader)
		line, isPrefix, err := rd.rd.ReadLine()
		if err != nil {
			return err
		}
		if isPrefix {
			return ErrLineTooLong
		}
		prefix := line[0]
		switch prefix {
		case ErrorReply:
			return errors.New(string(line[1:]))
		case ArrayReply:
			size, err := parseInt(line[1:])
			if err != nil {
				return err
			}
			s.v = make([]string, size)
			for i := 0; i < size; i++ {
				p, err := rd.Parse()
				if err != nil {
					return err
				}
				s.v[i] = string(p)
			}
			return nil
		default:
			return ErrInvalidReply
		}
	}
}

func (s *ScanParser) Scan() (cursor uint64, keys []string, err error) {
	if s.err != nil {
		return 0, nil, s.err
	}
	if len(s.v) == 0 {
		return 0, nil, nil
	}
	cursor, err = strconv.ParseUint(s.v[0], 10, 64)
	if err != nil {
		return 0, nil, err
	}
	return cursor, s.v[1:], nil
}

func (s *ScanParser) ScanResult() *ScanResult {
	cursor, keys, err := s.Scan()
	return &ScanResult{
		cursor: cursor,
		val:    keys,
		err:    err,
	}
}

func (s *ScanParser) SetErr(err error) {
	s.err = err
	if err != nil {
		s.v = nil
		s.cursor = 0
	}
}

type KeyValueParser struct {
	key   string
	value string
	err   error
}

func (k *KeyValueParser) Parse(ctx context.Context, reader io.Reader) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		rd := NewProtocolParser(reader)
		line, isPrefix, err := rd.rd.ReadLine()
		if err != nil {
			return err
		}
		if isPrefix {
			return ErrLineTooLong
		}
		prefix := line[0]
		switch prefix {
		case ErrorReply:
			return errors.New(string(line[1:]))
		case ArrayReply:
			size, err := parseInt(line[1:])
			if err != nil {
				return err
			}
			if size != 2 {
				return ErrInvalidReply
			}
			key, err := rd.Parse()
			if err != nil {
				return err
			}
			value, err := rd.Parse()
			if err != nil {
				return err
			}
			k.key = string(key)
			k.value = string(value)
			return nil
		default:
			return ErrInvalidReply
		}
	}
}

func (k *KeyValueParser) SetErr(err error) {
	k.err = err
	if err != nil {
		k.key = ""
		k.value = ""
	}
}

func (k *KeyValueParser) KeyValueResult() *KeyValueResult {
	return &KeyValueResult{
		key:   k.key,
		value: k.value,
		err:   k.err,
	}
}
