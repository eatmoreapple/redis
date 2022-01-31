package redis

import "errors"

// Error represents an error from the Redis server.
type Error struct {
	err error
}

// Error implements the error interface.
func (e *Error) Error() string {
	return e.err.Error()
}

// make sure Error implements the error interface
var _ error = (*Error)(nil)

// NewError returns an Error from the given error.
func NewError(err error) error {
	return &Error{err: err}
}

// ErrNil indicates that a reply value is nil.
var (
	ErrNil           error = NewError(errors.New("redis: nil returned"))
	InvalidTypeError error = errors.New("redis: invalid type")
	PrefixError      error = errors.New("redis: line prefix error")
)

// IsNilError returns true if the error is ErrNil.
func IsNilError(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, ErrNil)
}
