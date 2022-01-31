package redis

type (
	LInsertOp string
	LOrderOp  string
)

const (
	Before LInsertOp = "BEFORE"
	After  LInsertOp = "AFTER"

	ASC  LOrderOp = "ASC"
	DESC LOrderOp = "DESC"
)

func (c *Conn) LPush(key string, value ...interface{}) (int64, error) {
	result := &IntResult{}
	args := []interface{}{"LPUSH", key}
	args = append(args, value...)
	err := c.Send(result, args...)
	return result.Int64(), err
}

func (c *Conn) RPush(key string, value ...interface{}) (int64, error) {
	result := &IntResult{}
	args := []interface{}{"RPUSH", key}
	args = append(args, value...)
	err := c.Send(result, args...)
	return result.Int64(), err
}

func (c *Conn) LRange(key string, start, end int64) ([]string, error) {
	result := &StringArrayResult{}
	err := c.Send(result, "LRANGE", key, start, end)
	return result.Strings(), err
}

func (c *Conn) LPop(key string) (string, error) {
	result := &StringResult{}
	err := c.Send(result, "LPOP", key)
	return result.String(), err
}

func (c *Conn) RPop(key string) (string, error) {
	result := &StringResult{}
	err := c.Send(result, "RPOP", key)
	return result.String(), err
}

func (c *Conn) LLen(key string) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "LLEN", key)
	return result.Int64(), err
}

func (c *Conn) LRem(key string, count int64, value interface{}) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "LREM", key, count, value)
	return result.Int64(), err
}

func (c *Conn) LIndex(key string, index int64) (string, error) {
	result := &StringResult{}
	err := c.Send(result, "LINDEX", key, index)
	return result.String(), err
}

func (c *Conn) LSet(key string, index int64, value interface{}) (bool, error) {
	result := &BoolResult{}
	err := c.Send(result, "LSET", key, index, value)
	return result.Bool(), err
}

func (c *Conn) LTrim(key string, start, end int64) (bool, error) {
	result := &BoolResult{}
	err := c.Send(result, "LTRIM", key, start, end)
	return result.Bool(), err
}

func (c *Conn) LInsert(key string, op LInsertOp, pivot, value interface{}) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "LINSERT", key, op, pivot, value)
	return result.Int64(), err
}

func (c *Conn) LInsertBefore(key string, pivot, value interface{}) (int64, error) {
	return c.LInsert(key, Before, pivot, value)
}

func (c *Conn) LInsertAfter(key string, pivot, value interface{}) (int64, error) {
	return c.LInsert(key, After, pivot, value)
}

func (c *Conn) RPopLPush(key string, dest string) (string, error) {
	result := &StringResult{}
	err := c.Send(result, "RPOPLPUSH", key, dest)
	return result.String(), err
}

func (c *Conn) Sort(key string) ([]string, error) {
	result := &StringArrayResult{}
	err := c.Send(result, "SORT", key)
	return result.Strings(), err
}

func (c *Conn) SortAlpha(key string) ([]string, error) {
	result := &StringArrayResult{}
	err := c.Send(result, "SORT", key, "ALPHA")
	return result.Strings(), err
}

func (c *Conn) SortOrder(key string, order LOrderOp) ([]string, error) {
	result := &StringArrayResult{}
	err := c.Send(result, "SORT", key, order)
	return result.Strings(), err
}

func (c *Conn) SortAsc(key string) ([]string, error) {
	return c.SortOrder(key, ASC)
}

func (c *Conn) SortDesc(key string) ([]string, error) {
	return c.SortOrder(key, DESC)
}

func (c *Conn) SortLimit(key string, offset, count int64) ([]string, error) {
	result := &StringArrayResult{}
	err := c.Send(result, "SORT", key, "LIMIT", offset, count)
	return result.Strings(), err
}
