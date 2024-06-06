package db

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"log"
)

type redisClinet struct {
	cli *redis.Client
}

func NewRedisClient(addr string) *redisClinet {
	return &redisClinet{
		redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}
}

func (c *redisClinet) Hmset(key string, fields map[string]interface{}) error {
	return c.cli.HMSet(key, fields).Err()
}

func (c *redisClinet) Hmget(key string, fields ...string) (map[string]string, error) {
	ans, err := c.cli.HMGet(key, fields...).Result()
	if err != nil {
		log.Fatalf("redis hmset error: %v", err)
		return nil, errors.New(err.Error())
	}
	res := make(map[string]string, len(fields))
	for i := 0; i < len(fields); i++ {
		res[fields[i]] = ans[i].(string)
	}
	return res, nil
}
