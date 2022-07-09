package redis

import (
	"context"
	"fmt"
)

// SubscribeMessage is a message received from the server.
type SubscribeMessage struct {
	Kind    string
	Channel string
	Message []byte
}

// String returns the string representation of the message.
func (s SubscribeMessage) String() string {
	return fmt.Sprintf("Kind: %s, Channel: %s, Message: %s", s.Kind, s.Channel, s.Message)
}

func (s SubscribeMessage) IsUnsubscribe() bool {
	return s.Kind == "unsubscribe"
}

// PSubscribeMessage is a message received from the server.
type PSubscribeMessage struct {
	Kind    string
	Pattern string
	Channel string
	Message []byte
}

// String returns the string representation of the message.
func (p PSubscribeMessage) String() string {
	return fmt.Sprintf("Kind: %s, Pattern: %s, Channel: %s, Message: %s", p.Kind, p.Pattern, p.Channel, p.Message)
}

func (p PSubscribeMessage) IsUnsubscribe() bool {
	return p.Kind == "punsubscribe"
}

// UnsubscribeMessage is a message received from the server.
type UnsubscribeMessage struct {
	Kind    string
	Channel string
	Count   int64
}

// String returns the string representation of the message.
func (u UnsubscribeMessage) String() string {
	return fmt.Sprintf("Kind: %s, Channel: %s, Count: %d", u.Kind, u.Channel, u.Count)
}

// Publish sends a message to a channel.
// The message should not be empty.
func (c *Conn) Publish(channel string, message interface{}) (int64, error) {
	result := &IntResult{}
	err := c.Send(result, "PUBLISH", channel, message)
	return result.Int64(), err
}

// Subscribe subscribes the client to the specified channels.
// The messages received from the server are pushed to the given channel.
// And it will be blocked until the client unsubscribes the patterns.
func (c *Conn) Subscribe(ch chan *SubscribeMessage, channels ...string) error {
	result := &SubscribeResult{ch: ch}
	args := append([]interface{}{"SUBSCRIBE"}, StringToInterface(channels...)...)
	err := c.Send(result, args...)
	return err
}

// UnSubscribe unsubscribes the client from the given channels.
func (c *Conn) UnSubscribe(channels ...string) (*UnsubscribeMessage, error) {
	result := &UnsubscribeResult{}
	args := append([]interface{}{"UNSUBSCRIBE"}, StringToInterface(channels...)...)
	err := c.Send(result, args...)
	return result.msg, err
}

// PSubscribe subscribes the client to the given patterns.
// The messages received from the server are pushed to the given channel.
// And it will be blocked until the client unsubscribes the patterns.
func (c *Conn) PSubscribe(ctx context.Context, ch chan *PSubscribeMessage, patterns ...string) error {
	result := &PSubscribeResult{ch: ch, ctx: ctx}
	args := append([]interface{}{"PSUBSCRIBE"}, StringToInterface(patterns...)...)
	err := c.Send(result, args...)
	return err
}

// PUnSubscribe unsubscribes the client from the given patterns.
func (c *Conn) PUnSubscribe(patterns ...string) (*UnsubscribeMessage, error) {
	result := &UnsubscribeResult{}
	args := append([]interface{}{"PUNSUBSCRIBE"}, StringToInterface(patterns...)...)
	err := c.Send(result, args...)
	return result.msg, err
}
