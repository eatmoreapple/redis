package redis

import (
	"errors"
	"strconv"
)

// Result represents a Redis command result.
type Result interface {
	Parse(reader *Reader) error
}

// StringResult represents a string result.
type StringResult struct{ v string }

// Parse implements the Result interface.
func (r *StringResult) Parse(reader *Reader) error {
	line, isPrefix, err := reader.rd.ReadLine()
	if err != nil {
		return err
	}
	if isPrefix {
		return PrefixError
	}
	prefix := line[0]
	switch prefix {
	case StatusReply:
		r.v = string(line[1:])
		return nil
	case ErrorReply:
		return NewError(errors.New(string(line[1:])))
	case StringReply:
		size, err := parseInt(line[1:])
		if err != nil {
			return err
		}
		if size == -1 {
			return ErrNil
		}
		p := make([]byte, size+2)
		_, err = reader.rd.Read(p)
		if err != nil {
			return err
		}
		r.v = string(p[:size])
		return nil
	}
	return InvalidTypeError
}

// String returns the string result.
func (r *StringResult) String() string {
	return r.v
}

// BoolResult represents a bool result.
type BoolResult struct{ v bool }

// Parse implements the Result interface.
func (b *BoolResult) Parse(reader *Reader) error {
	line, isPrefix, err := reader.rd.ReadLine()
	if err != nil {
		return err
	}
	if isPrefix {
		return PrefixError
	}
	prefix := line[0]
	switch prefix {
	case StatusReply:
		b.v = string(line[1:]) == "OK"
		return nil
	case ErrorReply:
		return NewError(errors.New(string(line[1:])))
	case IntReply:
		b.v = string(line[1:]) != "0"
		return nil
	}
	return InvalidTypeError
}

// Bool returns the bool result.
func (b *BoolResult) Bool() bool {
	return b.v
}

// IntResult represents an int result.
type IntResult struct{ v int64 }

// Parse implements the Result interface.
func (i *IntResult) Parse(reader *Reader) error {
	line, isPrefix, err := reader.rd.ReadLine()
	if err != nil {
		return err
	}
	if isPrefix {
		return PrefixError
	}
	prefix := line[0]
	switch prefix {
	case StatusReply:
		v, err := parseInt(line[1:])
		if err != nil {
			return err
		}
		i.v = int64(v)
		return err
	case ErrorReply:
		return NewError(errors.New(string(line[1:])))
	case IntReply:
		v, err := parseInt(line[1:])
		if err != nil {
			return err
		}
		i.v = int64(v)
		return err
	}
	return InvalidTypeError
}

// Int64 returns the int result.
func (i *IntResult) Int64() int64 {
	return i.v
}

// MapResult represents a map result.
type MapResult struct{ v H }

// Parse implements the Result interface.
func (m *MapResult) Parse(reader *Reader) error {
	line, isPrefix, err := reader.rd.ReadLine()
	if err != nil {
		return err
	}
	if isPrefix {
		return PrefixError
	}
	prefix := line[0]
	switch prefix {
	case ErrorReply:
		return NewError(errors.New(string(line[1:])))
	case ArrayReply:
		size, err := parseInt(line[1:])
		if err != nil {
			return err
		}
		if size%2 != 0 {
			return errors.New("invalid reply: map must have even number of elements")
		}
		m.v = make(map[string]string, size/2)
		var key string
		for i := 0; i < size; i++ {
			p, err := reader.Parse()
			if err != nil {
				return err
			}
			if len(key) == 0 {
				key = string(p)
			} else {
				m.v[key] = string(p)
				key = ""
			}
		}
		return nil
	}
	return InvalidTypeError
}

// StringMap returns the map result.
func (m *MapResult) StringMap() H {
	return m.v
}

// StringArrayResult represents a string array result.
type StringArrayResult struct{ v []string }

// Parse implements the Result interface.
func (r *StringArrayResult) Parse(reader *Reader) error {
	line, isPrefix, err := reader.rd.ReadLine()
	if err != nil {
		return err
	}
	if isPrefix {
		return PrefixError
	}
	prefix := line[0]
	switch prefix {
	case ErrorReply:
		return NewError(errors.New(string(line[1:])))
	case ArrayReply:
		size, err := parseInt(line[1:])
		if err != nil {
			return err
		}
		r.v = make([]string, size)
		for i := 0; i < size; i++ {
			p, err := reader.Parse()
			if err != nil {
				return err
			}
			r.v[i] = string(p)
		}
		return nil
	}
	return InvalidTypeError
}

// Strings returns the string array result.
func (r *StringArrayResult) Strings() []string {
	return r.v
}

// FloatResult represents a float result.
type FloatResult struct{ v float64 }

// Parse implements the Result interface.
func (f *FloatResult) Parse(reader *Reader) error {
	line, isPrefix, err := reader.rd.ReadLine()
	if err != nil {
		return err
	}
	if isPrefix {
		return PrefixError
	}
	prefix := line[0]
	switch prefix {
	case StatusReply:
		v, err := parseInt(line[1:])
		if err != nil {
			return err
		}
		f.v = float64(v)
		return err
	case ErrorReply:
		return NewError(errors.New(string(line[1:])))
	case StringReply:
		size, err := parseInt(line[1:])
		if err != nil {
			return err
		}
		p := make([]byte, size+2)
		_, err = reader.rd.Read(p)
		if err != nil {
			return err
		}
		v := string(p[:size])
		f.v, err = strconv.ParseFloat(v, 64)
		return err
	}
	return InvalidTypeError

}

// Float64 returns the float result.
func (f *FloatResult) Float64() float64 {
	return f.v
}
