package redis

import (
	"context"

	"github.com/Yamon955/ShortVideo/user/entity/conf"
	"github.com/redis/go-redis/v9"
	"trpc.group/trpc-go/trpc-go/log"
)

var redisCli *redis.Client

// Init 初始化 redis 连接对象
func Init() error {
	redisCli = redis.NewClient(&redis.Options{
		Addr:     conf.GetAppConfig().RedisConf.Address,
		Password: conf.GetAppConfig().RedisConf.Password,
		DB:       conf.GetAppConfig().RedisConf.DB,
	})
	// 测试连接
	_, err := redisCli.Ping(context.Background()).Result()
	if err != nil {
		log.Errorf("ping redis failed, redis conf:%+v, err:%v", conf.GetAppConfig().RedisConf, err)
		return err
	}
	return nil
}

// GetClient 获取连接对象
func GetClient() *redis.Client {
	return redisCli
}
