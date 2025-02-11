package redis

import (
	"context"
)

type RDBClient interface {
	// BitAnd bitmap and operation
	BitAnd(ctx context.Context, key1, key2 string) int64
	// ZRangeByScore zrange
	ZRangeByScore(ctx context.Context, key string, min, max int64) []string
}

func NewRDBClient() RDBClient {
	return &rdbClient{
		client: redisCli,
	}
}
