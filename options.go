package redis

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

var dataSourceRegx = regexp.MustCompile(`:(?P<password>.*?)@(?P<network>.*)\((?P<address>.*):(?P<port>\d+)\)/(?P<db>\d+)`)

type Options struct {
	Password        string
	NetWork         string
	Address         string
	Port            int
	DB              int
	MaxOpenNums     int
	ConnMaxLifetime time.Duration
}

func (o Options) address() string {
	return o.Address + ":" + strconv.Itoa(o.Port)
}

func NewOptionsWithDataSource(dataSource string) (*Options, error) {
	subMatch := dataSourceRegx.FindStringSubmatch(dataSource)
	if len(subMatch) != 6 {
		return nil, errors.New("redis: invalid data source")
	}
	port, err := strconv.Atoi(subMatch[4])
	if err != nil {
		return nil, errors.New("redis: invalid data source")
	}
	if subMatch[5] == "" {
		subMatch[5] = "0"
	}
	db, err := strconv.Atoi(subMatch[5])
	if err != nil {
		return nil, errors.New("redis: invalid data source")
	}
	return &Options{
		Password: subMatch[1],
		NetWork:  subMatch[2],
		Address:  subMatch[3],
		Port:     port,
		DB:       db,
	}, nil
}
