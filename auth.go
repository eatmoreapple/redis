// Copyright 2022 eatmoreapple.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redis

import (
	"context"
	"time"
)

func (c *Client) Auth(ctx context.Context, password string) *BoolResult {
	args := Args{"AUTH", password}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) Echo(ctx context.Context, message string) *StringResult {
	args := Args{"ECHO", message}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringResult()
}

func (c *Client) Ping(ctx context.Context) *StringResult {
	args := Args{"PING"}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringResult()
}

func (c *Client) Quit(ctx context.Context) *StringResult {
	args := Args{"QUIT"}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringResult()
}

func (c *Client) Select(ctx context.Context, index int) *StringResult {
	args := Args{"SELECT", index}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringResult()
}

func (c *Client) SwapDB(ctx context.Context, index1, index2 int) *StringResult {
	args := Args{"SWAPDB", index1, index2}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringResult()
}

func (c *Client) Wait(ctx context.Context, numSlaves int, timeout time.Duration) *IntResult {
	args := Args{"WAIT", numSlaves, timeout}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) Unlink(ctx context.Context, keys ...string) *IntResult {
	args := Args{"UNLINK"}.AddFlat(keys)
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) UnlinkCtx(ctx context.Context, keys ...string) *IntResult {
	args := Args{"UNLINK"}.AddFlat(keys)
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) Del(ctx context.Context, keys ...string) *IntResult {
	args := Args{"DEL"}.AddFlat(keys)
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) DelCtx(ctx context.Context, keys ...string) *IntResult {
	args := Args{"DEL"}.AddFlat(keys)
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) Dump(ctx context.Context, key string) *StringResult {
	args := Args{"DUMP", key}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringResult()
}

func (c *Client) Exists(ctx context.Context, keys ...string) *IntResult {
	args := Args{"EXISTS"}.AddFlat(keys)
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) ExistsCtx(ctx context.Context, keys ...string) *IntResult {
	args := Args{"EXISTS"}.AddFlat(keys)
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) Expire(ctx context.Context, key string, expiration time.Duration) *BoolResult {
	args := Args{"EXPIRE", key, expiration / time.Second}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) ExpireAt(ctx context.Context, key string, tm time.Time) *BoolResult {
	args := Args{"EXPIREAT", key, tm.Unix()}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) Keys(ctx context.Context, pattern string) *StringSliceResult {
	args := Args{"KEYS", pattern}
	var parser StringSliceParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringSliceResult()
}

func (c *Client) Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *BoolResult {
	args := Args{"MIGRATE", host, port, key, db, timeout / time.Millisecond}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) Move(ctx context.Context, key string, db int) *BoolResult {
	args := Args{"MOVE", key, db}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) ObjectRefCount(ctx context.Context, key string) *IntResult {
	args := Args{"OBJECT", "REFCOUNT", key}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.IntResult()
}

func (c *Client) ObjectEncoding(ctx context.Context, key string) *StringResult {
	args := Args{"OBJECT", "ENCODING", key}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringResult()
}

func (c *Client) ObjectIdleTime(ctx context.Context, key string) *DurationResult {
	args := Args{"OBJECT", "IDLETIME", key}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.DurationResult()
}

func (c *Client) Persist(ctx context.Context, key string) *BoolResult {
	args := Args{"PERSIST", key}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) PExpire(ctx context.Context, key string, expiration time.Duration) *BoolResult {
	args := Args{"PEXPIRE", key, expiration / time.Millisecond}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) PExpireAt(ctx context.Context, key string, tm time.Time) *BoolResult {
	args := Args{"PEXPIREAT", key, tm.UnixNano() / int64(time.Millisecond)}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) PTTL(ctx context.Context, key string) *DurationResult {
	args := Args{"PTTL", key}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.DurationResult()
}

func (c *Client) RandomKey(ctx context.Context) *StringResult {
	args := Args{"RANDOMKEY"}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.StringResult()
}

func (c *Client) Rename(ctx context.Context, key, newKey string) *BoolResult {
	args := Args{"RENAME", key, newKey}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) RenameNX(ctx context.Context, key, newKey string) *BoolResult {
	args := Args{"RENAMENX", key, newKey}
	var parser IntegersParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) Restore(ctx context.Context, key string, ttl time.Duration, value string) *BoolResult {
	args := Args{"RESTORE", key, ttl / time.Millisecond, value}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}

func (c *Client) RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *BoolResult {
	args := Args{"RESTORE", key, ttl / time.Millisecond, value, "REPLACE"}
	var parser StringParser
	err := c.Call(ctx, args, &parser)
	parser.SetErr(err)
	return parser.BoolResult()
}
