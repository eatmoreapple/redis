package redis

import (
	"encoding"
	"errors"
	"io"
	"reflect"
	"strconv"
	"strings"
)

const (
	ErrorReply  = '-'
	StatusReply = '+'
	IntReply    = ':'
	StringReply = '$'
	ArrayReply  = '*'
)

type Builder interface {
	WriteArgs(args ...interface{}) error
	io.WriterTo
}

type builder struct {
	bd strings.Builder
}

func (w *builder) writePrefix(n int) error {
	var err error
	if _, err = w.bd.Write([]byte{ArrayReply}); err != nil {
		return err
	}
	if _, err = w.bd.WriteString(strconv.Itoa(n)); err != nil {
		return err
	}
	if _, err = w.bd.Write([]byte{'\r', '\n'}); err != nil {
		return err
	}
	return err
}

func (w *builder) WriteArgs(args ...interface{}) error {
	var err error
	if err = w.writePrefix(len(args)); err != nil {
		return err
	}
	for _, arg := range args {
		if err = w.writeArg(arg); err != nil {
			return err
		}
	}
	return err
}

func (w *builder) String() string {
	return w.bd.String()
}

func (w *builder) WriteTo(writer io.Writer) (int64, error) {
	n, err := writer.Write([]byte(w.String()))
	return int64(n), err
}

func (w *builder) writeArg(v interface{}) error {
	switch t := v.(type) {
	case nil:
		return w.writeLine("")
	case string:
		return w.writeLine(t)
	case []byte:
		return w.writeLine(string(t))
	case int8:
		return w.writeInt(int64(t))
	case int16:
		return w.writeInt(int64(t))
	case int32:
		return w.writeInt(int64(t))
	case int64:
		return w.writeInt(t)
	case int:
		return w.writeInt(int64(t))
	case uint8:
		return w.writeInt(int64(t))
	case uint16:
		return w.writeInt(int64(t))
	case uint32:
		return w.writeInt(int64(t))
	case uint64:
		return w.writeInt(int64(t))
	case uint:
		return w.writeInt(int64(t))
	case float32:
		return w.writeFloat(float64(t))
	case float64:
		return w.writeFloat(t)
	case bool:
		return w.writeBool(t)
	case encoding.BinaryMarshaler:
		return w.writeBinaryMarshaler(t)
	}
	return errors.New("redis: can't marshal " + reflect.TypeOf(v).String())
}

func (w *builder) writeLine(content string) error {
	var err error
	if _, err = w.bd.Write([]byte{StringReply}); err != nil {
		return err
	}
	if _, err = w.bd.WriteString(strconv.Itoa(len(content))); err != nil {
		return err
	}
	if _, err = w.bd.Write([]byte{'\r', '\n'}); err != nil {
		return err
	}
	if _, err = w.bd.WriteString(content); err != nil {
		return err
	}
	if _, err = w.bd.Write([]byte{'\r', '\n'}); err != nil {
		return err
	}
	return err
}

func (w *builder) writeInt(i int64) error {
	return w.writeLine(strconv.FormatInt(i, 10))
}

func (w *builder) writeFloat(i float64) error {
	return w.writeLine(strconv.FormatFloat(i, 'f', -1, 64))
}

func (w *builder) writeBool(i bool) error {
	if i {
		return w.writeLine("1")
	}
	return w.writeLine("0")
}

func (w *builder) writeBinaryMarshaler(Marshaler encoding.BinaryMarshaler) error {
	data, err := Marshaler.MarshalBinary()
	if err != nil {
		return err
	}
	return w.writeLine(string(data))
}
