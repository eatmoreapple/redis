package redis

import (
	"bufio"
	"net"
)

type Conn struct {
	conn   net.Conn
	client *Client
}

func (c *Conn) Close() error {
	return c.conn.Close()
}

func (c *Conn) Release() {
	c.client.Release(c)
}

func (c *Conn) Do(builder Builder, result Result, args ...interface{}) error {
	if err := builder.WriteArgs(args...); err != nil {
		return err
	}
	if _, err := builder.WriteTo(c.conn); err != nil {
		return err
	}
	reader := &Reader{rd: bufio.NewReader(c.conn)}
	return result.Parse(reader)
}

func (c *Conn) Send(result Result, args ...interface{}) error {
	bd := &builder{}
	return c.Do(bd, result, args...)
}

func (c *Conn) Ping() (string, error) {
	result := &StringResult{}
	err := c.Send(result, "PING")
	return result.String(), err
}

func (c *Conn) Keys(pattern string) ([]string, error) {
	result := &StringArrayResult{}
	err := c.Send(result, "KEYS", pattern)
	return result.Strings(), err
}

func (c *Conn) Select(index int) (bool, error) {
	result := &BoolResult{}
	err := c.Send(result, "SELECT", index)
	return result.Bool(), err
}

func (c *Conn) Move(key string, index int) (bool, error) {
	result := &BoolResult{}
	err := c.Send(result, "MOVE", key, index)
	return result.Bool(), err
}

func (c *Conn) Flush(index int) (bool, error) {
	result := &BoolResult{}
	err := c.Send(result, "FLUSH", index)
	return result.Bool(), err
}

func (c *Conn) Del(keys ...string) (bool, error) {
	result := &BoolResult{}
	args := StringToInterface(keys...)
	args = append([]interface{}{"DEL"}, args...)
	err := c.Send(result, args...)
	return result.Bool(), err
}

func (c *Conn) Type(key string) (string, error) {
	result := &StringResult{}
	err := c.Send(result, "TYPE", key)
	return result.String(), err
}

func (c *Conn) RandomKey() (string, error) {
	result := &StringResult{}
	err := c.Send(result, "RANDOMKEY")
	return result.String(), err
}

func (c *Conn) Exists(key string) (bool, error) {
	result := &BoolResult{}
	err := c.Send(result, "EXISTS", key)
	return result.Bool(), err
}

func (c *Conn) Expire(key string, seconds int) (bool, error) {
	result := &BoolResult{}
	err := c.Send(result, "EXPIRE", key, seconds)
	return result.Bool(), err
}

func (c *Conn) PExpire(key string, milliseconds int) (bool, error) {
	result := &BoolResult{}
	err := c.Send(result, "PEXPIRE", key, milliseconds)
	return result.Bool(), err
}

func (c *Conn) Persist(key string) (bool, error) {
	result := &BoolResult{}
	err := c.Send(result, "PERSIST", key)
	return result.Bool(), err
}
