// Copyright 2022 eatmoreapple.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redis

import "errors"

var (
	// ErrInvalidReply is returned when the reply is not valid.
	ErrInvalidReply = errors.New("invalid reply type")

	// ErrLineTooLong is returned when the reply line is too long.
	ErrLineTooLong = errors.New("line too long")

	// ErrNil indicates that a reply value is nil.
	ErrNil = errors.New("nil returned")

	// ErrBadConn indicates that the netConn is bad.
	ErrBadConn = errors.New("bad netConn")

	// ErrPoolTimeout indicates that the pool is timeout.
	ErrPoolTimeout = errors.New("pool timeout")
)
