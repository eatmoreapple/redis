package redis

func (c *Conn) SAdd(key string, members ...interface{}) (int64, error) {
	result := &IntResult{}
	args := append([]interface{}{"SADD", key}, members...)
	err := c.Send(result, args...)
	return result.Int64(), err
}

func (c *Conn) SMembers(key string) ([]string, error) {
	result := &StringArrayResult{}
	err := c.Send(result, "SMEMBERS", key)
	return result.Strings(), err
}

func (c *Conn) SRem(key string, members ...interface{}) (int64, error) {
	result := &IntResult{}
	args := append([]interface{}{"SREM", key}, members...)
	err := c.Send(result, args...)
	return result.Int64(), err
}

func (c *Conn) SIsMember(key string, member interface{}) (bool, error) {
	result := &BoolResult{}
	err := c.Send(result, "SISMEMBER", key, member)
	return result.Bool(), err
}

func (c *Conn) SCard(key string) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "SCARD", key)
	return result.Int64(), err
}

func (c *Conn) SDiff(keys ...string) ([]string, error) {
	result := &StringArrayResult{}
	args := append([]interface{}{"SDIFF"}, StringToInterface(keys...)...)
	err := c.Send(result, args...)
	return result.Strings(), err
}

func (c *Conn) SInter(keys ...string) ([]string, error) {
	result := &StringArrayResult{}
	args := append([]interface{}{"SINTER"}, StringToInterface(keys...)...)
	err := c.Send(result, args...)
	return result.Strings(), err
}

func (c *Conn) SUnion(keys ...string) ([]string, error) {
	result := &StringArrayResult{}
	args := append([]interface{}{"SUNION"}, StringToInterface(keys...)...)
	err := c.Send(result, args...)
	return result.Strings(), err
}

func (c *Conn) SRandMember(key string, count ...int64) ([]string, error) {
	result := &StringArrayResult{}
	args := []interface{}{"SRANDMEMBER", key}
	if len(count) > 0 {
		args = append(args, count[0])
	}
	err := c.Send(result, args...)
	return result.Strings(), err
}

func (c *Conn) SPop(key string, count ...int64) ([]string, error) {
	result := &StringArrayResult{}
	args := []interface{}{"SPOP", key}
	if len(count) > 0 {
		args = append(args, count[0])
	}
	err := c.Send(result, args...)
	return result.Strings(), err
}
