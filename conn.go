// Copyright 2022 eatmoreapple.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redis

import (
	"context"
	"io"
	"net"
	"time"
)

// netConn is a wrapper of net.Conn.
type netConn interface {
	net.Conn

	// Context returns the context of the netConn.
	Context() context.Context

	// LastUsedTime returns the last used time of the netConn.
	LastUsedTime() time.Time

	// Active reset the last used time of the netConn.
	Active()

	// Err returns the error of the netConn.
	Err() error

	Call(reader io.Reader, writer io.Writer) error
}

var _ netConn = (*connection)(nil)

// badCoon wraps a bad connection.
var badCoon = &connection{err: ErrBadConn}

type connection struct {
	net.Conn
	ctx          context.Context
	cancel       context.CancelFunc
	createTime   time.Time
	lastUsedTime time.Time
	err          error
}

func (c *connection) Call(reader io.Reader, writer io.Writer) error {
	if _, err := io.Copy(c.Conn, reader); err != nil {
		return err
	}
	if _, err := io.Copy(writer, c.Conn); err != nil {
		return err
	}
	return nil
}

func (c *connection) LastUsedTime() time.Time {
	if c.lastUsedTime.IsZero() {
		return c.createTime
	}
	return c.lastUsedTime
}

func (c *connection) Active() {
	c.lastUsedTime = time.Now()
}

func (c *connection) Err() error {
	return c.err
}

func (c *connection) Context() context.Context {
	if c.ctx == nil {
		return context.Background()
	}
	return c.ctx
}

func (c *connection) AliveTime() time.Duration {
	return time.Since(c.createTime)
}

// Close closes the netConn.
func (c *connection) Close() error {
	if c.cancel != nil {
		c.cancel()
	}
	return c.Conn.Close()
}

func (c *connection) auth(username, password string) error {
	// if the password is set, authenticate the connection
	args := Args{"AUTH"}
	// if the username is set, authenticate the connection with username
	if len(username) > 0 {
		args = AppendArg(args, username)
	}
	// authenticate the connection with password
	args = AppendArg(args, password)
	// wrap the connection with a reader and a writer
	reader := newArgsWrapper(c.ctx, args)
	var parser StringParser
	// wrap the connection with a reader and a writer
	writer := newParserWrapper(c.ctx, &parser)
	err := c.Call(reader, writer)
	// set the error of the connection
	parser.SetErr(err)
	return parser.BoolResult().Err()
}

func (c *Client) conn(ctx context.Context) (conn netConn, err error) {
	// check if pool timeout is set
	if c.opt.PoolTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, c.opt.PoolTimeout)
		defer cancel()
	}
LOOP:
	for {
		c.signalCH <- incrSignal
		select {
		case <-c.ctx.Done():
			return nil, c.ctx.Err()
		case <-ctx.Done():
			return nil, ctx.Err()
		case conn = <-c.connCH:
			// check if the connection is valid
			if err = conn.Err(); err != nil {
				return nil, err
			}

			// check if the connection is expired
			if c.opt.IdleTimeout > 0 && time.Since(conn.LastUsedTime()) > c.opt.IdleTimeout {
				goto CLOSE
			}
			select {
			case <-conn.Context().Done():
				goto CLOSE
			default:
				// active the connection
				conn.Active()
				return conn, nil
			}
		}
		// CLOSE is a label for closing the connection
	CLOSE:
		// call the OnClose callback
		if c.opt.OnClose != nil {
			c.opt.OnClose(c.ctx, conn)
		}

		// close the connection
		_ = conn.Close()

		// send a signal to the listener
		c.signalCH <- decrSignal

		// retry to get a connection
		continue LOOP
	}
}

func (c *Client) release(conn netConn) {
	c.connCH <- conn
}

func (c *Client) signalListener() {
	var currentConnNum int32
	for {
		select {
		case <-c.ctx.Done():
			return
		case signal := <-c.signalCH:
			if !signal {
				currentConnNum--
				break
			}
			if currentConnNum >= c.opt.PoolSize {
				break
			}
			var (
				conn net.Conn
				err  error
			)
			// try to connect to the redis server
			for i := 0; i < c.opt.MaxRetries; i++ {
				conn, err = c.opt.Dialer(c.ctx, c.opt.Network, c.opt.Addr)
				if err == nil && conn != nil {
					break
				}
			}

			// if failed, return
			if err != nil {
				c.connCH <- badCoon
				break
			}

			if c.opt.OnConnect != nil {
				c.opt.OnConnect(c.ctx, conn)
			}

			cn := &connection{
				Conn:       conn,
				createTime: time.Now(),
			}

			if c.opt.MaxConnAge > 0 {
				// if the deadline is not set, use a context
				cn.ctx, cn.cancel = context.WithDeadline(c.ctx, time.Now().Add(c.opt.MaxConnAge))
			}

			// send the connection to the connection channel
			if len(c.opt.Password) > 0 {
				// try to authenticate the connection
				if err = cn.auth(c.opt.Username, c.opt.Password); err != nil {
					c.connCH <- &connection{err: err}
					break
				}
			}

			c.connCH <- cn
			currentConnNum++
		}
	}
}

type signalKind bool

const (
	incrSignal signalKind = true
	decrSignal signalKind = false
)
