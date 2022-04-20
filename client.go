package redis

import (
	"errors"
	"fmt"
	"net"
	"sync"
)

var ErrPoolClosed = errors.New("redis: get on closed pool")

func Dial(dataSource string) (*Client, error) {
	options, err := NewOptionsWithDataSource(dataSource)
	if err != nil {
		return nil, err
	}
	return &Client{options: options}, nil
}

type Client struct {
	closed   bool
	openNums int
	freeNums int
	connList chan *Conn
	lock     sync.RWMutex
	options  *Options
}

// Get a connection from the pool.
// If no connections are available, it will create a new one.
// If maxOpen is reached, it will wait for an available connection.
func (p *Client) Get() (*Conn, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.closed {
		return nil, ErrPoolClosed
	}
	// init pool if needed
	if p.connList == nil {
		p.connList = make(chan *Conn, p.options.MaxOpenNums)
	}
	// if no connection limit or connection limit not reached, create a new connection
	if p.openNums == 0 || p.openNums < p.options.MaxOpenNums {
		conn, err := p.newConn()
		if err != nil {
			return nil, err
		}
		// increase openNums
		p.openNums++
		return conn, nil
	} else {
		// get a connection from the pool
		conn := <-p.connList
		p.freeNums--
		return conn, nil
	}
}

func (p *Client) newConn() (*Conn, error) {
	// create a new connection
	c, err := net.Dial(p.options.NetWork, p.options.address())
	if err != nil {
		return nil, err
	}
	coon := &Conn{pool: p, conn: c}
	if err := coon.init(p.options); err != nil {
		return nil, err
	}
	return coon, nil
}

func (p *Client) Close() error {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.closed {
		return nil
	}
	close(p.connList)
	p.closed = true
	var errorList []error
	for c := range p.connList {
		if err := c.Close(); err != nil {
			errorList = append(errorList, err)
		}
	}
	if len(errorList) > 0 {
		return fmt.Errorf("redis: close error: %v", errorList)
	}
	return nil
}

func (p *Client) release(c *Conn) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	// if pool is closed or no open connection limit, close the connection
	if p.closed || p.options.MaxOpenNums <= 0 {
		return c.Close()
	}
	// if current connection is expired, close it
	if c.isExpired() {
		p.openNums--
		return c.Close()
	} else {
		// otherwise, put it back to the pool
		p.connList <- c
		p.freeNums++
		return nil
	}
}

func (p *Client) SetMaxOpenNums(i int) {
	p.options.MaxOpenNums = i
}
