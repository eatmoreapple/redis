// Copyright 2022 eatmoreapple.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redis

import (
	"context"
	"time"
)

func (c *Client) Set(ctx context.Context, key string, value any, expireIn time.Duration) *BoolResult {
	args := Args{"SET", key, value}
	if seconds := int64(expireIn.Seconds()); seconds > 0 {
		args = append(args, "EX", seconds)
	}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) Get(ctx context.Context, key string) *StringResult {
	args := Args{"GET", key}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringResult()
}

func (c *Client) GetRange(ctx context.Context, key string, start, end int) *StringResult {
	args := Args{"GETRANGE", key, start, end}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringResult()
}

func (c *Client) GetSet(ctx context.Context, key string, value any) *StringResult {
	args := Args{"GETSET", key, value}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringResult()
}

func (c *Client) SetRange(ctx context.Context, key string, offset int, value any) *IntResult {
	args := Args{"SETRANGE", key, offset, value}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) GetBit(ctx context.Context, key string, offset int) *IntResult {
	args := Args{"GETBIT", key, offset}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) SetBit(ctx context.Context, key string, offset int, value int64) *IntResult {
	args := Args{"SETBIT", key, offset, value}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) BitCount(ctx context.Context, key string, startAndEnd ...int64) *IntResult {
	args := Args{"BITCOUNT", key}
	if len(startAndEnd) == 2 {
		args = append(args, startAndEnd[0], startAndEnd[1])
	}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) MGet(ctx context.Context, key string, keys ...string) *StringSliceResult {
	var parser StringSliceParser
	args := Args{"MGET", key}
	args = AppendArg(Args{"MGET"}, keys...)
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringSliceResult()
}

func (c *Client) MSet(ctx context.Context, key string, value any, keysValues ...any) *BoolResult {
	var parser StringParser
	args := Args{"MSET", key, value}
	args = AppendArg(args, keysValues...)
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) MSetNX(ctx context.Context, key string, value any, keysValues ...any) *BoolResult {
	var parser IntegersParser
	args := Args{"MSETNX", key, value}
	args = AppendArg(args, keysValues...)
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) Incr(ctx context.Context, key string) *IntResult {
	args := Args{"INCR", key}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) IncrBy(ctx context.Context, key string, increment int64) *IntResult {
	args := Args{"INCRBY", key, increment}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) IncrByFloat(ctx context.Context, key string, increment float64) *FloatResult {
	args := Args{"INCRBYFLOAT", key, increment}
	var parser FloatsParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.FloatResult()
}

func (c *Client) Decr(ctx context.Context, key string) *IntResult {
	args := Args{"DECR", key}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) DecrBy(ctx context.Context, key string, decrement int64) *IntResult {
	args := Args{"DECRBY", key, decrement}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) Append(ctx context.Context, key string, value any) *IntResult {
	args := Args{"APPEND", key, value}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) StrLen(ctx context.Context, key string) *IntResult {
	args := Args{"STRLEN", key}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) SetNX(ctx context.Context, key string, value any) *BoolResult {
	args := Args{"SETNX", key, value}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) SetXX(ctx context.Context, key string, value any) *BoolResult {
	args := Args{"SET", key, value, "XX"}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) SetEX(ctx context.Context, key string, value any, expiration time.Duration) *BoolResult {
	args := Args{"SETEX", key, expiration / time.Second, value}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) PSetEX(ctx context.Context, key string, value any, expiration time.Duration) *BoolResult {
	args := Args{"PSETEX", key, expiration / time.Millisecond, value}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}
