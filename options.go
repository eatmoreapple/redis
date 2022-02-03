package redis

import "time"

type Options struct {
	NetWork     string
	Addr        string
	Password    string
	DB          int
	MaxLifetime time.Duration
	MaxOpenNums int
	WaitTimeout time.Duration
}

func (o *Options) init() {
	if o.NetWork == "" {
		o.NetWork = "tcp"
	}
	if o.Addr == "" {
		o.Addr = "localhost:6379"
	}
	if o.MaxOpenNums == 0 {
		o.MaxOpenNums = 1
	}
}
