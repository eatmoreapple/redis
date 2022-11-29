// Copyright 2022 eatmoreapple.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redis

import (
	"context"
	"net"
	"runtime"
	"time"
)

type Options struct {
	// Addr is the TCP address of the Redis server.
	// The address must be in the form "host:port".
	// If no port is specified, then 6379 is used.
	// If no host is specified, then "localhost" is used.
	// To use a Unix domain socket, provide a path to a socket file.
	Addr string
	// Network is the type of network to use.
	// It must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket".
	// Default is "tcp".
	Network string
	// Dialer is an optional dialer for creating TCP connections.
	// If Dialer is nil, then net.Dialer is used.
	Dialer func(ctx context.Context, network, addr string) (net.Conn, error)
	// OnConnect is an optional hook executed when a new connection is established.
	// It receives a context and a connection. It is safe for OnConnect to read and
	// write to the connection.
	OnConnect func(ctx context.Context, cn net.Conn)
	// OnClose is an optional hook executed when a connection is closed.
	// It receives a context and a connection. It is safe for OnClose to read and
	// write to the connection.
	OnClose func(ctx context.Context, cn net.Conn)
	// Username is the username for AUTH (requires Redis 6 or above).
	Username string
	// Password is the password for AUTH (requires Redis 6 or above).
	Password string
	// DB is the Redis database to select.
	DB int
	// MaxRetries specifies the maximum number of retries before giving up.
	// Default is 3 retries.
	MaxRetries int
	// PoolSize is the maximum number of socket connections.
	// Default is 2 * runtime.NumCPU.
	PoolSize int32
	// MaxConnAge is the maximum amount of time a connection may be reused.
	// Expired connections may be closed lazily before reuse.
	// If the value is zero, then connections are reused forever.
	MaxConnAge time.Duration
	// PoolTimeout is the amount of time client waits for connection if all
	// connections are busy before returning an error.
	// If the value is zero and there are no idle connections, then
	// ErrPoolTimeout is returned.
	PoolTimeout time.Duration
	// IdleTimeout is the maximum amount of time an idle connection may remain idle
	// before closing itself.
	// If the value is zero, then idle connections are not closed.
	IdleTimeout time.Duration
}

func (o *Options) init() {
	if o.Addr == "" {
		o.Addr = "localhost:6379"
	}
	if o.Network == "" {
		o.Network = "tcp"
	}
	if o.Dialer == nil {
		o.Dialer = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return net.Dial(network, addr)
		}
	}
	if o.PoolSize == 0 {
		o.PoolSize = int32(runtime.NumCPU()) * 2
	}
	if o.MaxRetries == 0 {
		o.MaxRetries = 3
	}
}

type OptionFunc func(*Options)

func (f OptionFunc) apply(o *Options) {
	f(o)
}

func WithAddr(addr string) OptionFunc {
	return OptionFunc(func(o *Options) {
		o.Addr = addr
	})
}

func WithNetwork(network string) OptionFunc {
	return OptionFunc(func(o *Options) {
		o.Network = network
	})
}

func WithDialer(dialer func(ctx context.Context, network, addr string) (net.Conn, error)) OptionFunc {
	return OptionFunc(func(o *Options) {
		o.Dialer = dialer
	})
}

func WithOnConnect(onConnect func(ctx context.Context, cn net.Conn)) OptionFunc {
	return OptionFunc(func(o *Options) {
		o.OnConnect = onConnect
	})
}

func WithOnClose(onClose func(ctx context.Context, cn net.Conn)) OptionFunc {
	return OptionFunc(func(o *Options) {
		o.OnClose = onClose
	})
}

func WithUsername(username string) OptionFunc {
	return OptionFunc(func(o *Options) {
		o.Username = username
	})
}

func WithPassword(password string) OptionFunc {
	return OptionFunc(func(o *Options) {
		o.Password = password
	})
}

func WithDB(db int) OptionFunc {
	return OptionFunc(func(o *Options) {
		o.DB = db
	})
}

func WithMaxRetries(maxRetries int) OptionFunc {
	return OptionFunc(func(o *Options) {
		o.MaxRetries = maxRetries
	})
}

func WithPoolSize(poolSize int) OptionFunc {
	return OptionFunc(func(o *Options) {
		o.PoolSize = int32(poolSize)
	})
}

func WithMaxConnAge(maxConnAge time.Duration) OptionFunc {
	return OptionFunc(func(o *Options) {
		o.MaxConnAge = maxConnAge
	})
}

func WithPoolTimeout(poolTimeout time.Duration) OptionFunc {
	return OptionFunc(func(o *Options) {
		o.PoolTimeout = poolTimeout
	})
}
