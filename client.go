package redis

import (
	"net"
)

type Client struct {
	pool           chan *Conn
	maxConnections int
	addr           string
	network        string
}

func NewClient(network, addr string, maxConnections int) (*Client, error) {
	c := &Client{
		pool:           make(chan *Conn, maxConnections),
		maxConnections: maxConnections,
		addr:           addr,
		network:        network,
	}
	for i := 0; i < maxConnections; i++ {
		conn, err := c.newConn()
		if err != nil {
			return nil, err
		}
		c.pool <- conn
	}
	return c, nil
}

func (c Client) Release(conn *Conn) {
	select {
	case c.pool <- conn:
	default:
		conn.Close()
	}
}

func (c *Client) newConn() (*Conn, error) {
	conn, err := net.Dial(c.network, c.addr)
	if err != nil {
		return nil, err
	}
	return &Conn{
		conn:   conn,
		client: c,
	}, nil
}

func (c *Client) Conn() (*Conn, error) {
	select {
	case conn := <-c.pool:
		return conn, nil
	default:
		return c.newConn()
	}
}
