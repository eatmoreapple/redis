// Copyright 2022 eatmoreapple.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redis

import (
	"bytes"
	"context"
	"encoding"
	"errors"
	"reflect"
	"strconv"
)

// Args defines a list of arguments for a command.
type Args []any

func (a Args) AddFlat(keys []string) Args {
	if len(keys) == 0 {
		return a
	}
	for _, key := range keys {
		a = append(a, key)
	}
	return a
}

func AppendArg[T any](dst Args, args ...T) Args {
	if len(args) == 0 {
		return dst
	}
	if dst == nil {
		dst = make(Args, 0, len(args))
	}
	for _, arg := range args {
		dst = append(dst, arg)
	}
	return dst
}

// ArgsBuilder is a builder for Args.
type ArgsBuilder struct {
	*bytes.Buffer
}

// Build builds args with resp protocol.
func (s *ArgsBuilder) Build(ctx context.Context, args Args) error {
	if len(args) == 0 {
		return errors.New("redis: empty command")
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if err := s.WriteArgs(args); err != nil {
			return err
		}
	}
	return nil
}

// WriteArgs writes args with resp protocol.
func (s *ArgsBuilder) WriteArgs(args Args) error {
	var err error
	if err = s.writePrefix(len(args)); err != nil {
		return err
	}
	for _, arg := range args {
		if err = s.writeArg(arg); err != nil {
			return err
		}
	}
	return err
}

func (s *ArgsBuilder) writeArg(v any) error {
	switch t := v.(type) {
	case nil:
		return s.writeLine("")
	case string:
		return s.writeLine(t)
	case []byte:
		return s.writeLine(string(t))
	case int8:
		return s.writeInt(int64(t))
	case int16:
		return s.writeInt(int64(t))
	case int32:
		return s.writeInt(int64(t))
	case int64:
		return s.writeInt(t)
	case int:
		return s.writeInt(int64(t))
	case uint8:
		return s.writeInt(int64(t))
	case uint16:
		return s.writeInt(int64(t))
	case uint32:
		return s.writeInt(int64(t))
	case uint64:
		return s.writeInt(int64(t))
	case uint:
		return s.writeInt(int64(t))
	case float32:
		return s.writeFloat(float64(t))
	case float64:
		return s.writeFloat(t)
	case bool:
		return s.writeBool(t)
	case encoding.BinaryMarshaler:
		return s.writeBinaryMarshaller(t)
	}
	return errors.New("redis: can't marshal " + reflect.TypeOf(v).String())
}

func (s *ArgsBuilder) writeLine(content string) error {
	var err error
	if _, err = s.Write([]byte{StringReply}); err != nil {
		return err
	}
	if _, err = s.WriteString(strconv.Itoa(len(content))); err != nil {
		return err
	}
	if _, err = s.Write(crlf); err != nil {
		return err
	}
	if _, err = s.WriteString(content); err != nil {
		return err
	}
	if _, err = s.Write(crlf); err != nil {
		return err
	}
	return err
}

func (s *ArgsBuilder) writeInt(i int64) error {
	return s.writeLine(strconv.FormatInt(i, 10))
}

func (s *ArgsBuilder) writeFloat(i float64) error {
	return s.writeLine(strconv.FormatFloat(i, 'f', -1, 64))
}

func (s *ArgsBuilder) writeBool(i bool) error {
	if i {
		return s.writeLine("1")
	}
	return s.writeLine("0")
}

func (s *ArgsBuilder) writeBinaryMarshaller(Marshaller encoding.BinaryMarshaler) error {
	data, err := Marshaller.MarshalBinary()
	if err != nil {
		return err
	}
	return s.writeLine(string(data))
}

func (s *ArgsBuilder) writePrefix(n int) error {
	var err error
	if _, err = s.Write([]byte{ArrayReply}); err != nil {
		return err
	}
	if _, err = s.WriteString(strconv.Itoa(n)); err != nil {
		return err
	}
	if _, err = s.Write(crlf); err != nil {
		return err
	}
	return err
}
