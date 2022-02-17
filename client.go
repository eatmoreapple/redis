package redis

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
)

var ErrPoolClosed = errors.New("redis: get on closed pool")

func NewClientWithOptions(options *Options) *Client {
	options.init()
	return &Client{
		options: options,
		conns:   make(chan *Conn, options.MaxOpenNums),
	}
}

func NewClient(opts ...*Options) *Client {
	opt := &Options{}
	if len(opts) > 0 {
		opt = opts[0]
	}
	return NewClientWithOptions(opt)
}

type Client struct {
	closed   bool
	openNums int
	freeNums int
	conns    chan *Conn
	lock     sync.RWMutex
	options  *Options
}

func (p *Client) Get() (*Conn, error) {
	p.lock.RLock()
	if p.closed {
		p.lock.RLocker()
		return nil, ErrPoolClosed
	}
	p.lock.RUnlock()
	if !p.hasFreeConns() {
		return p.newConn()
	}
	p.lock.Lock()
	p.freeNums--
	p.lock.Unlock()
	return <-p.conns, nil
}

func (p *Client) newConn() (*Conn, error) {
	p.lock.Lock()
	if p.openNums >= p.options.MaxOpenNums {
		p.lock.Unlock()
		if p.options.WaitTimeout > 0 {
			ctx, cancel := context.WithTimeout(context.Background(), p.options.WaitTimeout)
			defer cancel()
			select {
			case conn := <-p.conns:
				return conn, nil
			case <-ctx.Done():
				return p.dailNewCoon()
			}
		} else {
			return <-p.conns, nil
		}
	}
	p.lock.Unlock()
	return p.dailNewCoon()
}

func (p *Client) dailNewCoon() (*Conn, error) {
	c, err := net.Dial(p.options.NetWork, p.options.Addr)
	if err != nil {
		return nil, err
	}
	coon := &Conn{pool: p, conn: c}
	if err := coon.init(p.options); err != nil {
		return nil, err
	}
	p.lock.Lock()
	if p.openNums < p.options.MaxOpenNums {
		p.openNums++
	}
	p.lock.Unlock()
	return coon, nil
}

func (p *Client) hasFreeConns() bool {
	return p.freeNums > 0
}

func (p *Client) Close() error {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.closed {
		return nil
	}
	close(p.conns)
	p.closed = true
	var errorList []error
	for c := range p.conns {
		if err := c.Close(); err != nil {
			errorList = append(errorList, err)
		}
	}
	return fmt.Errorf("%v", errorList)
}

func (p *Client) isFull() bool {
	return len(p.conns) >= p.options.MaxOpenNums
}

func (p *Client) release(c *Conn) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.closed || c.isExpired() || p.isFull() {
		return c.Close()
	}
	p.conns <- c
	p.freeNums++
	return nil
}
