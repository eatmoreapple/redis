package redis

func (c *Conn) Set(key string, value interface{}) (bool, error) {
	result := &BoolResult{}
	err := c.Send(result, "SET", key, value)
	return result.Bool(), err
}

func (c *Conn) Get(key string) (string, error) {
	result := &StringResult{}
	err := c.Send(result, "GET", key)
	return result.String(), err
}

func (c *Conn) GetRange(key string, start, end int) (string, error) {
	result := &StringResult{}
	err := c.Send(result, "GETRANGE", key, start, end)
	return result.String(), err
}

func (c *Conn) GetSet(key string, value interface{}) (string, error) {
	result := &StringResult{}
	err := c.Send(result, "GETSET", key, value)
	return result.String(), err
}

func (c *Conn) GetBit(key string, offset int) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "GETBIT", key, offset)
	return result.Int64(), err
}

func (c *Conn) SetBit(key string, offset int, value int64) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "SETBIT", key, offset, value)
	return result.Int64(), err
}

func (c *Conn) BitCount(key string, startAndEnd ...int64) (int64, error) {
	result := &IntResult{}
	args := []interface{}{"BITCOUNT", key}
	if len(startAndEnd) == 2 {
		args = append(args, startAndEnd[0], startAndEnd[1])
	}
	err := c.Send(result, args...)
	return result.Int64(), err
}

func (c *Conn) MGet(key ...string) ([]string, error) {
	result := &StringArrayResult{}
	err := c.Send(result, "MGET", key)
	return result.Strings(), err
}

func (c *Conn) SetEx(key string, seconds int, value interface{}) (bool, error) {
	result := &BoolResult{}
	err := c.Send(result, "SETEX", key, seconds, value)
	return result.Bool(), err
}

func (c *Conn) SetNx(key string, value interface{}) (bool, error) {
	result := &BoolResult{}
	err := c.Send(result, "SETNX", key, value)
	return result.Bool(), err
}

func (c *Conn) SetRange(key string, offset int, value interface{}) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "SETRANGE", key, offset, value)
	return result.Int64(), err
}

func (c *Conn) StrLen(key string) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "STRLEN", key)
	return result.Int64(), err
}

func (c *Conn) MSet(vs Q) (bool, error) {
	result := &BoolResult{}
	args := MapToInterface("MSET", vs)
	err := c.Send(result, args...)
	return result.Bool(), err
}

func (c *Conn) MSetNx(vs Q) (bool, error) {
	result := &BoolResult{}
	args := MapToInterface("MSETNX", vs)
	err := c.Send(result, args...)
	return result.Bool(), err
}

func (c *Conn) PsSetEx(key string, milliseconds int, value interface{}) (bool, error) {
	result := &BoolResult{}
	err := c.Send(result, "PSETEX", key, milliseconds, value)
	return result.Bool(), err
}

func (c *Conn) Incr(key string) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "INCR", key)
	return result.Int64(), err
}

func (c *Conn) IncrBy(key string, value int64) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "INCRBY", key, value)
	return result.Int64(), err
}

func (c *Conn) IncrByFloat(key string, value float64) (float64, error) {
	result := &FloatResult{}
	err := c.Send(result, "INCRBYFLOAT", key, value)
	return result.Float64(), err
}

func (c *Conn) Decr(key string) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "DECR", key)
	return result.Int64(), err
}

func (c *Conn) DecrBy(key string, value int64) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "DECRBY", key, value)
	return result.Int64(), err
}

func (c *Conn) Append(key string, value string) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "APPEND", key, value)
	return result.Int64(), err
}
