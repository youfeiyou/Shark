package db

import (
	"context"
	//"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"log"
)

type redisClinet struct {
	cli *redis.Client
}

var MsgSeqScript = redis.NewScript(`
   local num_arg = #ARGV
   local res = {}
   for i = 1, num_arg do
      local v = redis.call("HINCRBY",KEYS[1], ARGV[i], 1)
      res[#res+1] = v
   end
   return redis.call("HGETALL",KEYS[1])
   
`)

func NewRedisClient(addr string) *redisClinet {
	return &redisClinet{
		redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: "shark",
		}),
	}
}

func (c *redisClinet) HMSet(key string, fields map[string]interface{}) error {
	return c.cli.HMSet(context.Background(), key, fields).Err()
}

func (c *redisClinet) HMGet(key string, fields ...string) (map[string]string, error) {
	ans, err := c.cli.HMGet(context.Background(), key, fields...).Result()
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
	var cmd *redis.IntCmd
	if cmd = c.cli.HIncrBy(context.Background(), key, field, incr); cmd.Err() != nil {
		log.Printf("reids api Hincrby fail: %+v", cmd.Err())
		return cmd.Err()
	}
	return nil
}

func (c *redisClinet) Eval(script string, keys []string, argv ...interface{}) (*redis.Cmd, error) {
	var cmd *redis.Cmd
	if cmd = c.cli.Eval(context.Background(), script, keys, argv); cmd.Err() != nil {
		log.Printf("reids api Hincrby fail: %+v", cmd.Err())
		return nil, cmd.Err()
	}
	return cmd, nil
}

func (c *redisClinet) Incr(key string) (uint64, error) {
	var cmd *redis.IntCmd
	if cmd = c.cli.Incr(context.Background(), key); cmd.Err() != nil {
		log.Printf("reids api Incr fail: %+v", cmd.Err())
		return 0, cmd.Err()
	}
	return uint64(cmd.Val()), nil
}

func (c *redisClinet) HDel(key string, fields ...string) error {
	var cmd *redis.IntCmd
	if cmd = c.cli.HDel(context.Background(), key, fields...); cmd.Err() != nil {
		log.Printf("reids api HDel fail: %+v", cmd.Err())
		return cmd.Err()
	}
	return nil
}
