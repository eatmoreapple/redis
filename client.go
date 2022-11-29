// Copyright 2022 eatmoreapple.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redis

import (
	"context"
	"io"
	"net"
)

func NewClient(ctx context.Context, opts ...OptionFunc) *Client {
	client := &Client{ctx: ctx}
	for _, opt := range opts {
		opt.apply(&client.opt)
	}
	client.init()
	return client
}

func NewClientWithOption(ctx context.Context, opt Options) *Client {
	client := &Client{ctx: ctx, opt: opt}
	client.init()
	return client
}

type Client struct {
	Dialer   func(ctx context.Context) (net.Conn, error)
	connCH   chan netConn
	signalCH chan signalKind
	ctx      context.Context
	cancel   context.CancelFunc
	opt      Options
}

func (c *Client) init() {
	c.opt.init()
	c.connCH = make(chan netConn, c.opt.PoolSize)
	c.signalCH = make(chan signalKind)
	c.ctx, c.cancel = context.WithCancel(c.ctx)
	go c.signalListener()
}

func (c *Client) Call(ctx context.Context, args Args, parser Parser) error {
	reader := newArgsWrapper(ctx, args)
	writer := newParserWrapper(ctx, parser)
	return c.call(ctx, reader, writer)
}

func (c *Client) call(ctx context.Context, reader io.Reader, writer io.Writer) error {
	// get connection from pool
	conn, err := c.conn(ctx)
	if err != nil {
		return err
	}
	defer c.release(conn)
	return conn.Call(reader, writer)
}

func (c *Client) Close() {
	defer c.cancel()
	close(c.signalCH)
	close(c.connCH)
	for conn := range c.connCH {
		_ = conn.Close()
	}
}
