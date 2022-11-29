// Copyright 2022 eatmoreapple.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redis

import (
	"context"
)

func (c *Client) HSet(ctx context.Context, key string, field string, value any) *BoolResult {
	args := Args{"HSET", key, field, value}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) HSetNX(ctx context.Context, key string, field string, value any) *BoolResult {
	args := Args{"HSETNX", key, field, value}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) HGet(ctx context.Context, key string, field string) *StringResult {
	args := Args{"HGET", key, field}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringResult()
}

func (c *Client) HGetAll(ctx context.Context, key string) *StringMapResult {
	args := Args{"HGETALL", key}
	var parser StringMapParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringMapResult()
}

func (c *Client) HDel(ctx context.Context, key string, fields ...string) *IntResult {
	args := Args{"HDEL", key}
	args = AppendArg(args, fields...)
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) HExists(ctx context.Context, key string, field string) *BoolResult {
	args := Args{"HEXISTS", key, field}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) HIncrBy(ctx context.Context, key string, field string, increment int) *IntResult {
	args := Args{"HINCRBY", key, field, increment}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) HIncrByFloat(ctx context.Context, key string, field string, increment float64) *FloatResult {
	args := Args{"HINCRBYFLOAT", key, field, increment}
	var parser FloatsParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.FloatResult()
}

func (c *Client) HKeys(ctx context.Context, key string) *StringSliceResult {
	args := Args{"HKEYS", key}
	var parser StringSliceParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringSliceResult()
}

func (c *Client) HVals(ctx context.Context, key string) *StringSliceResult {
	args := Args{"HVALS", key}
	var parser StringSliceParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringSliceResult()
}

func (c *Client) HLen(ctx context.Context, key string) *IntResult {
	args := Args{"HLEN", key}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) HMGet(ctx context.Context, key string, fields ...string) *StringSliceResult {
	args := Args{"HMGET", key}
	args = AppendArg(args, fields...)
	var parser StringSliceParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringSliceResult()
}

func (c *Client) HMSet(ctx context.Context, key string, values map[string]string) *BoolResult {
	args := Args{"HMSET", key}
	for k, v := range values {
		args = AppendArg(args, k, v)
	}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) HScan(ctx context.Context, key string, cursor int, match string, count int) *ScanResult {
	args := Args{"HSCAN", key, cursor}
	if match != "" {
		args = AppendArg(args, "MATCH", match)
	}
	if count > 0 {
		args = append(args, "COUNT", count)
	}
	var parser ScanParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.ScanResult()
}

func (c *Client) HStrLen(ctx context.Context, key string, field string) *IntResult {
	args := Args{"HSTRLEN", key, field}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) HSetMulti(ctx context.Context, key string, values map[string]string) *BoolResult {
	args := Args{"HMSET", key}
	for k, v := range values {
		args = AppendArg(args, k, v)
	}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) HGetMulti(ctx context.Context, key string, fields ...string) *StringMapResult {
	args := Args{"HMGET", key}
	args = AppendArg(args, fields...)
	var parser StringMapParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringMapResult()
}

func (c *Client) HGetAllMap(ctx context.Context, key string) *StringMapResult {
	args := Args{"HGETALL", key}
	var parser StringMapParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringMapResult()
}
