package redis

func (c *Conn) ZAdd(key string, score float64, member interface{}) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "ZADD", key, score, member)
	return result.Int64(), err
}

func (c *Conn) ZIncrBy(key string, increment float64, member interface{}) (float64, error) {
	result := &FloatResult{}
	err := c.Send(result, "ZINCRBY", key, increment, member)
	return result.Float64(), err
}

func (c *Conn) ZScore(key string, member interface{}) (float64, error) {
	result := &FloatResult{}
	err := c.Send(result, "ZSCORE", key, member)
	return result.Float64(), err
}

func (c *Conn) ZRange(key string, start, stop int) ([]string, error) {
	result := &StringArrayResult{}
	err := c.Send(result, "ZRANGE", key, start, stop)
	return result.Strings(), err
}

func (c *Conn) ZRangeWithScores(key string, start, stop int) (H, error) {
	result := &MapResult{}
	err := c.Send(result, "ZRANGE", key, start, stop, "WITHSCORES")
	return result.StringMap(), err
}

func (c *Conn) ZRangeRev(key string, start, stop int) ([]string, error) {
	result := &StringArrayResult{}
	err := c.Send(result, "ZRANGE", key, start, stop, "REV")
	return result.Strings(), err
}

func (c *Conn) ZRangeByScore(key string, start, stop float64) ([]string, error) {
	result := &StringArrayResult{}
	err := c.Send(result, "ZRANGE", key, start, stop, "BYSCORE")
	return result.Strings(), err
}

func (c *Conn) ZRangeByScoreLimit(key string, start, stop float64, offset, count uint) ([]string, error) {
	result := &StringArrayResult{}
	err := c.Send(result, "ZRANGE", key, start, stop, "LIMIT", offset, count)
	return result.Strings(), err
}

func (c *Conn) ZCard(key string) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "ZCARD", key)
	return result.Int64(), err
}

func (c *Conn) ZCount(key string, min, max float64) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "ZCOUNT", key, min, max)
	return result.Int64(), err
}

func (c *Conn) ZRem(key string, members ...interface{}) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "ZREM", key, members)
	return result.Int64(), err
}

func (c *Conn) ZRemRangeByRank(key string, start, stop int) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "ZREMRANGEBYRANK", key, start, stop)
	return result.Int64(), err
}

func (c *Conn) ZRemRangeByLex(key string, min, max string) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "ZREMRANGEBYLEX", key, min, max)
	return result.Int64(), err
}

func (c *Conn) ZRemRangeByScore(key string, min, max float64) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "ZREMRANGEBYSCORE", key, min, max)
	return result.Int64(), err
}

func (c *Conn) ZRank(key string, member interface{}) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "ZRANK", key, member)
	return result.Int64(), err
}

func (c *Conn) ZRevRank(key string, member interface{}) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "ZREVRANK", key, member)
	return result.Int64(), err
}
