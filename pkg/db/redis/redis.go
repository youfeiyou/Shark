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

func (c *redisClinet) HMSet(key string, fields map[string]interface{}) error {
	return c.cli.HMSet(key, fields).Err()
}

func (c *redisClinet) HMGet(key string, fields ...string) (map[string]string, error) {
	ans, err := c.cli.HMGet(key, fields...).Result()
	if err != nil {
		log.Fatalf("redis hmset error: %v", err)
		return nil, errors.New(err.Error())
	}
	if len(ans) == 0 {
		return nil, nil
	}
	res := make(map[string]string, len(fields))
	for i := 0; i < len(fields); i++ {
		ok := false
		res[fields[i]], ok = ans[i].(string)
		if !ok {
			log.Printf("no value %+v:%+v", key, fields[i])
		}
	}
	return res, nil
}

func (c *redisClinet) HIncrBy(key, field string, incr int64) error {
	if err := c.cli.HIncrBy(key, field, incr); err != nil {
		log.Printf("reids api Hincrby fail: %+v", err)
		return err.Err()
	}
	return nil
}

func (c *redisClinet) Eval(script string, keys []string, argv ...interface{}) error {
	if err := c.cli.Eval(script, keys, argv); err != nil {
		log.Printf("reids api Hincrby fail: %+v", err)
		return err.Err()
	}
	return nil
}
