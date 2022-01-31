package redis

func (c *Conn) HSet(key string, vs Q) (int64, error) {
	result := &IntResult{}
	args := MapToInterface(key, vs)
	args = append([]interface{}{"HSET"}, args...)
	err := c.Send(result, args...)
	return result.Int64(), err
}

func (c *Conn) HGet(key string, field string) (string, error) {
	result := &StringResult{}
	err := c.Send(result, "HGET", key, field)
	return result.String(), err
}

func (c *Conn) HMSet(key string, vs Q) (string, error) {
	result := &StringResult{}
	args := MapToInterface(key, vs)
	args = append([]interface{}{"HMSET"}, args...)
	err := c.Send(result, args...)
	return result.String(), err
}

func (c *Conn) HMGet(key string, fields ...string) ([]string, error) {
	result := &StringArrayResult{}
	err := c.Send(result, "HMGET", key, fields)
	return result.Strings(), err
}

func (c *Conn) HGetAll(key string) (H, error) {
	result := &MapResult{}
	err := c.Send(result, "HGETALL", key)
	return result.StringMap(), err
}

func (c *Conn) HExists(key string, field string) (bool, error) {
	result := &BoolResult{}
	err := c.Send(result, "HEXISTS", key, field)
	return result.Bool(), err
}

func (c *Conn) HSetNx(key string, field string, value interface{}) (bool, error) {
	result := &BoolResult{}
	err := c.Send(result, "HSETNX", key, field, value)
	return result.Bool(), err
}

func (c *Conn) HIncrBy(key string, field string, value int64) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "HINCRBY", key, field, value)
	return result.Int64(), err
}

func (c *Conn) HDel(key string, fields ...string) (int64, error) {
	result := &IntResult{}
	args := append([]interface{}{"HDEL", key}, StringToInterface(fields...)...)
	err := c.Send(result, args...)
	return result.Int64(), err
}

func (c *Conn) HKeys(key string) ([]string, error) {
	result := &StringArrayResult{}
	err := c.Send(result, "HKEYS", key)
	return result.Strings(), err
}

func (c *Conn) HVals(key string) ([]string, error) {
	result := &StringArrayResult{}
	err := c.Send(result, "HVALS", key)
	return result.Strings(), err
}

func (c *Conn) HLen(key string) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "HLEN", key)
	return result.Int64(), err
}
