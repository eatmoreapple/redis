package redis

import (
	"bufio"
	"errors"
	"strconv"
)

type Reader struct {
	rd *bufio.Reader
}

func (r *Reader) Parse() ([]byte, error) {
	line, isPrefix, err := r.rd.ReadLine()
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
		p := make([]byte, size+2)
		_, err = r.rd.Read(p)
		if err != nil {
			return nil, err
		}
		return p[:size], nil
	case ArrayReply:
		size, err := parseInt(line[1:])
		if err != nil {
			return nil, err
		}
		var buf []byte
		for i := 0; i < size; i++ {
			p, err := r.Parse()
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
