// Copyright 2022 eatmoreapple.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redis

import (
	"context"
)

func (c *Client) SAdd(ctx context.Context, key string, member any, members ...any) *IntResult {
	args := Args{"SADD", key, member}
	args = AppendArg(args, members...)
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) SRem(ctx context.Context, key string, member any, members ...any) *IntResult {
	args := Args{"SREM", key, member}
	args = AppendArg(args, members...)
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) SPop(ctx context.Context, key string) *StringResult {
	args := Args{"SPOP", key}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringResult()
}

func (c *Client) SMove(ctx context.Context, source, destination string, member any) *BoolResult {
	args := Args{"SMOVE", source, destination, member}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) SCard(ctx context.Context, key string) *IntResult {
	args := Args{"SCARD", key}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) SIsMember(ctx context.Context, key string, member any) *BoolResult {
	args := Args{"SISMEMBER", key, member}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) SInter(ctx context.Context, keys ...string) *StringSliceResult {
	args := Args{"SINTER"}
	args = AppendArg(args, keys...)
	var parser StringSliceParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringSliceResult()
}

func (c *Client) SInterStore(ctx context.Context, destination string, keys ...string) *IntResult {
	args := Args{"SINTERSTORE", destination}
	args = AppendArg(args, keys...)
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) SUnion(ctx context.Context, keys ...string) *StringSliceResult {
	args := Args{"SUNION"}
	args = AppendArg(args, keys...)
	var parser StringSliceParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringSliceResult()
}

func (c *Client) SUnionStore(ctx context.Context, destination string, keys ...string) *IntResult {
	args := Args{"SUNIONSTORE", destination}
	args = AppendArg(args, keys...)
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) SDiff(ctx context.Context, keys ...string) *StringSliceResult {
	args := Args{"SDIFF"}
	args = AppendArg(args, keys...)
	var parser StringSliceParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringSliceResult()
}

func (c *Client) SDiffStore(ctx context.Context, destination string, keys ...string) *IntResult {
	args := Args{"SDIFFSTORE", destination}
	args = AppendArg(args, keys...)
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) SMembers(ctx context.Context, key string) *StringSliceResult {
	args := Args{"SMEMBERS", key}
	var parser StringSliceParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringSliceResult()
}

func (c *Client) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanResult {
	args := Args{"SSCAN", key, cursor}
	if match != "" {
		args = append(args, "MATCH", match)
	}
	if count > 0 {
		args = append(args, "COUNT", count)
	}
	var parser ScanParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.ScanResult()
}
