// Copyright 2022 eatmoreapple.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redis

import (
	"context"
	"time"
)

func (c *Client) LPush(ctx context.Context, key string, values ...any) *IntResult {
	args := Args{"LPUSH", key}
	args = AppendArg(args, values...)
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) LPushX(ctx context.Context, key string, values ...any) *IntResult {
	args := Args{"LPUSHX", key}
	args = AppendArg(args, values...)
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) RPush(ctx context.Context, key string, values ...any) *IntResult {
	args := Args{"RPUSH", key}
	args = AppendArg(args, values...)
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) RPushX(ctx context.Context, key string, values ...any) *IntResult {
	args := Args{"RPUSHX", key}
	args = AppendArg(args, values...)
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) LPop(ctx context.Context, key string) *StringResult {
	args := Args{"LPOP", key}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringResult()
}

func (c *Client) RPop(ctx context.Context, key string) *StringResult {
	args := Args{"RPOP", key}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringResult()
}

func (c *Client) LLen(ctx context.Context, key string) *IntResult {
	args := Args{"LLEN", key}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) LIndex(ctx context.Context, key string, index int) *StringResult {
	args := Args{"LINDEX", key, index}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringResult()
}

func (c *Client) LSet(ctx context.Context, key string, index int, value any) *BoolResult {
	args := Args{"LSET", key, index, value}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) LRange(ctx context.Context, key string, start, stop int) *StringSliceResult {
	args := Args{"LRANGE", key, start, stop}
	var parser StringSliceParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringSliceResult()
}

func (c *Client) LTrim(ctx context.Context, key string, start, stop int) *BoolResult {
	args := Args{"LTRIM", key, start, stop}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) LRem(ctx context.Context, key string, count int, value any) *IntResult {
	args := Args{"LREM", key, count, value}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) BLPop(ctx context.Context, keys []string, timeout time.Duration) *KeyValueResult {
	args := Args{"BLPOP"}
	args = AppendArg(args, keys...)
	args = AppendArg(args, int(timeout/time.Second))
	var parser KeyValueParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.KeyValueResult()
}

func (c *Client) BRPop(ctx context.Context, keys []string, timeout time.Duration) *KeyValueResult {
	args := Args{"BRPOP"}
	args = AppendArg(args, keys...)
	args = AppendArg(args, int(timeout/time.Second))
	var parser KeyValueParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.KeyValueResult()
}

func (c *Client) BRPopLPush(ctx context.Context, source string, destination string, timeout int) *StringResult {
	args := Args{"BRPOPLPUSH", source, destination, timeout}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringResult()
}

func (c *Client) RPopLPush(ctx context.Context, source, destination string) *StringResult {
	args := Args{"RPOPLPUSH", source, destination}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringResult()
}

func (c *Client) LInsertBefore(ctx context.Context, key string, pivot any, value any) *IntResult {
	args := Args{"LINSERT", key, "BEFORE", pivot, value}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) LInsertAfter(ctx context.Context, key string, pivot any, value any) *IntResult {
	args := Args{"LINSERT", key, "AFTER", pivot, value}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) LInsert(ctx context.Context, key string, pivot any, value any) *IntResult {
	return c.LInsertBefore(ctx, key, pivot, value)
}
