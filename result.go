// Copyright 2022 eatmoreapple.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redis

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type StringResult struct {
	val string
	err error
}

func (s StringResult) Result() (string, error) {
	return s.val, s.err
}

func (s StringResult) String() string {
	return s.val
}

func (s StringResult) Err() error {
	return s.err
}

func (s StringResult) Int() (int, error) {
	if s.err != nil {
		return 0, s.err
	}
	return strconv.Atoi(s.val)
}

func (s StringResult) Int64() (int64, error) {
	if s.err != nil {
		return 0, s.err
	}
	return strconv.ParseInt(s.val, 10, 64)
}

func (s StringResult) Uint64() (uint64, error) {
	if s.err != nil {
		return 0, s.err
	}
	return strconv.ParseUint(s.val, 10, 64)
}

func (s StringResult) Float64() (float64, error) {
	if s.err != nil {
		return 0, s.err
	}
	return strconv.ParseFloat(s.val, 64)
}

func (s StringResult) Bool() (bool, error) {
	return strconv.ParseBool(s.val)
}

func (s StringResult) Bytes() ([]byte, error) {
	return []byte(s.val), s.err
}

// MarshalJSON implements the json.Marshaler interface.
func (s StringResult) MarshalJSON() ([]byte, error) {
	return []byte(s.val), s.err
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s StringResult) UnmarshalJSON(data []byte) error {
	s.val = string(data)
	return nil
}

func (s StringResult) Scan(v any) error {
	if s.Err() != nil {
		return s.Err()
	}
	if scanner, ok := v.(Scanner); ok {
		return scanner.Scan(s.String())
	}
	switch v.(type) {
	case *string:
		*v.(*string) = s.String()
	case *[]byte:
		*v.(*[]byte) = []byte(s.String())
	case *int:
		i, err := s.Int()
		if err != nil {
			return err
		}
		*v.(*int) = i
	case *int64:
		i, err := s.Int64()
		if err != nil {
			return err
		}
		*v.(*int64) = i
	case *float64:
		f, err := s.Float64()
		if err != nil {
			return err
		}
		*v.(*float64) = f
	case *bool:
		b, err := s.Bool()
		if err != nil {
			return err
		}
		*v.(*bool) = b
	case *time.Time:
		t, err := time.Parse(time.RFC3339, s.String())
		if err != nil {
			return err
		}
		*v.(*time.Time) = t
	default:
		return fmt.Errorf("redis: can't scan into %T", v)
	}
	return nil
}

type BoolResult struct {
	val bool
	err error
}

func (b BoolResult) Result() (bool, error) {
	return b.val, b.err
}

func (b BoolResult) Bool() bool {
	return b.val
}

func (b BoolResult) Err() error {
	return b.err
}

func (b BoolResult) String() string {
	if b.val {
		return "true"
	}
	return "false"
}

type IntResult struct {
	val int64
	err error
}

func (i IntResult) Result() (int64, error) {
	return i.val, i.err
}

func (i IntResult) Int64() int64 {
	return i.val
}

func (i IntResult) Err() error {
	return i.err
}

func (i IntResult) String() string {
	return strconv.FormatInt(i.val, 10)
}

type StringSliceResult struct {
	val []string
	err error
}

func (s StringSliceResult) Result() ([]string, error) {
	return s.val, s.err
}

func (s StringSliceResult) Strings() []string {
	return s.val
}

func (s StringSliceResult) Err() error {
	return s.err
}

func (s StringSliceResult) Scan(v any) error {
	if s.Err() != nil {
		return s.Err()
	}
	if scanner, ok := v.(Scanner); ok {
		return scanner.Scan(s.Strings())
	}
	switch v.(type) {
	case *[]string:
		*v.(*[]string) = s.Strings()
	case *[]interface{}:
		vals := make([]interface{}, len(s.val))
		for i, v := range s.val {
			vals[i] = v
		}
		*v.(*[]interface{}) = vals
	default:
		return fmt.Errorf("redis: can't scan into %T", v)
	}
	return nil
}

type FloatResult struct {
	val float64
	err error
}

func (f FloatResult) Result() (float64, error) {
	return f.val, f.err
}

func (f FloatResult) Float64() float64 {
	return f.val
}

func (f FloatResult) Err() error {
	return f.err
}

func (f FloatResult) String() string {
	return strconv.FormatFloat(f.val, 'f', -1, 64)
}

func (f FloatResult) Scan(v any) error {
	if f.Err() != nil {
		return f.Err()
	}
	if scanner, ok := v.(Scanner); ok {
		return scanner.Scan(f.Float64())
	}
	switch v.(type) {
	case *float64:
		*v.(*float64) = f.Float64()
	case *float32:
		*v.(*float32) = float32(f.Float64())
	case *string:
		*v.(*string) = f.String()
	default:
		return fmt.Errorf("redis: can't scan into %T", v)
	}
	return nil
}

type StringMapResult struct {
	val map[string]string
	err error
}

func (s StringMapResult) Result() (map[string]string, error) {
	return s.val, s.err
}

func (s StringMapResult) StringMap() map[string]string {
	return s.val
}

func (s StringMapResult) Err() error {
	return s.err
}

func (s StringMapResult) Scan(v any) error {
	if s.Err() != nil {
		return s.Err()
	}
	if scanner, ok := v.(Scanner); ok {
		return scanner.Scan(s.StringMap())
	}
	switch v.(type) {
	case *map[string]string:
		*v.(*map[string]string) = s.StringMap()
	case *map[string]interface{}:
		vals := make(map[string]interface{}, len(s.val))
		for k, v := range s.val {
			vals[k] = v
		}
	default:
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Ptr {
			return errors.New("redis: Scan(non-pointer stringmap)")
		}
		if rv.IsNil() {
			return errors.New("redis: Scan(nil)")
		}
		rv = reflect.Indirect(rv)
		switch rv.Kind() {
		case reflect.Map:
			// get map element type
			elemType := rv.Type().Elem()
			switch elemType.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				vals := make(map[string]int64, len(s.val))
				for k, v := range s.val {
					val, err := strconv.ParseInt(v, 10, 64)
					if err != nil {
						return err
					}
					vals[k] = val
				}
			}

		}
		return fmt.Errorf("redis: can't scan into %T", v)
	}
	return nil
}

type ScanResult struct {
	cursor uint64
	val    []string
	err    error
}

func (s ScanResult) Result() (uint64, []string, error) {
	return s.cursor, s.val, s.err
}

func (s ScanResult) Cursor() uint64 {
	return s.cursor
}

func (s ScanResult) Values() []string {
	return s.val
}

func (s ScanResult) Err() error {
	return s.err
}

type DurationResult struct {
	val time.Duration
	err error
}

func (d DurationResult) Result() (time.Duration, error) {
	return d.val, d.err
}

func (d DurationResult) Duration() time.Duration {
	return d.val
}

func (d DurationResult) Err() error {
	return d.err
}

type KeyValueResult struct {
	key   string
	value string
	err   error
}

func (kv KeyValueResult) Key() string {
	return kv.key
}

func (kv KeyValueResult) Value() string {
	return kv.value
}

func (kv KeyValueResult) Err() error {
	return kv.err
}

func (kv KeyValueResult) Result() (string, string, error) {
	return kv.key, kv.value, kv.err
}

func (kv KeyValueResult) String() string {
	if kv.err != nil {
		return kv.err.Error()
	}
	return fmt.Sprintf("%s=%s", kv.key, kv.value)
}
