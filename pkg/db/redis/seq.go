package db

func GenSeq(key string) (uint64, error) {
	cli := NewRedisClient(RedisAddr)
	ret, err := cli.Incr(key)
	if err != nil {
		return 0, err
	}
	return ret, nil
}
