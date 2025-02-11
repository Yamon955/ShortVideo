package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type rdbClient struct {
	client *redis.Client
}

// BitAnd bitmap and operation
func (r *rdbClient) BitAnd(ctx context.Context, key1, key2 string) int64 {
	res := r.client.BitOpAnd(ctx, key1, key2).Val()
	return res
}

// ZRangeByScore zrange
func (r *rdbClient) ZRangeByScore(ctx context.Context, key string, min, max int64) []string {
	vids := r.client.ZRangeByScore().Val()
	return vids
}
