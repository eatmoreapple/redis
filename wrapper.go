// Copyright 2022 eatmoreapple.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redis

import (
	"context"
	"io"
)

// parserWrapper is a wrapper for Parser that implements io.Writer and io.ReaderFrom.
// It Hides the implementation details of Parser and Hook the io.Copy.
type parserWrapper struct {
	parser Parser
	ctx    context.Context
}

// Write implements io.Writer and call nothing.
func (p parserWrapper) Write(b []byte) (int, error) {
	return len(b), nil
}

// ReadFrom implements io.ReaderFrom and calls parser.Parse.
func (p parserWrapper) ReadFrom(r io.Reader) (int64, error) {
	return 0, p.parser.Parse(p.ctx, r)
}

// newParserWrapper returns a new parserWrapper.
func newParserWrapper(ctx context.Context, parser Parser) io.Writer {
	return &parserWrapper{ctx: ctx, parser: parser}
}

// argsWrapper is a wrapper for Args that implements io.Writer and io.Reader.
// It Hides the implementation details of Args and Hook the io.Copy.
type argsWrapper struct {
	args Args
	ctx  context.Context
}

// Read just implements io.Reader and call nothing.
func (a argsWrapper) Read(p []byte) (int, error) {
	return len(p), nil
}

// WriteTo implements io.WriterTo and calls args.WriteTo.
func (a argsWrapper) WriteTo(w io.Writer) (int64, error) {
	var buffer = getBuffer()
	defer putBuffer(buffer)
	builder := &ArgsBuilder{Buffer: buffer}
	if err := builder.Build(a.ctx, a.args); err != nil {
		return 0, err
	}
	return io.Copy(w, builder)
}

// newArgsWrapper returns a new argsWrapper.
func newArgsWrapper(ctx context.Context, args Args) io.Reader {
	return &argsWrapper{ctx: ctx, args: args}
}
