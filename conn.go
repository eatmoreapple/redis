package redis

import (
	"bufio"
	"net"
	"time"
)

type Conn struct {
	conn        net.Conn
	pool        *Client
	initialized bool
	createdAt   time.Time
}

func (c *Conn) init(options *Options) error {
	c.createdAt = time.Now()
	if c.initialized {
		return nil
	}

	if options.Password != "" {
		if _, err := c.Auth(options.Password); err != nil {
			return err
		}
	}

	if options.DB != 0 {
		if _, err := c.Select(options.DB); err != nil {
			return err
		}
	}

	c.initialized = true
	return nil
}

func (c *Conn) isExpired() bool {
	if c.pool.options.MaxLifetime == 0 {
		return false
	}
	return time.Now().Sub(c.createdAt) > c.pool.options.MaxLifetime
}

func (c *Conn) Close() error {
	return c.conn.Close()
}

func (c *Conn) Release() error {
	return c.pool.release(c)
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

func (c *Conn) FlushDB(op ...FlushOP) (bool, error) {
	result := &BoolResult{}
	var err error
	if len(op) == 0 {
		err = c.Send(result, "FLUSHDB")
	} else {
		err = c.Send(result, "FLUSHDB", op[0].String())
	}
	return result.Bool(), err
}

func (c *Conn) FlushDBAsync() (bool, error) {
	return c.FlushDB(ASYNC)
}

func (c *Conn) FlushDBSync() (bool, error) {
	return c.FlushDB(SYNC)
}

func (c *Conn) FlushAll(op ...FlushOP) (bool, error) {
	result := &BoolResult{}
	var err error
	if len(op) == 0 {
		err = c.Send(result, "FLUSHALL")
	} else {
		err = c.Send(result, "FLUSHALL", op[0].String())
	}
	return result.Bool(), err
}

func (c *Conn) FlushAllAsync() (bool, error) {
	return c.FlushAll(ASYNC)
}

func (c *Conn) FlushSync() (bool, error) {
	return c.FlushAll(SYNC)
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

func (c *Conn) Auth(password string) (bool, error) {
	result := &BoolResult{}
	err := c.Send(result, "AUTH", password)
	return result.Bool(), err
}
