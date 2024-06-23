package util

import (
	"math"
	"math/rand"
	db "shark/pkg/db/redis"
	"strconv"
)

const (
	shardingNum        = 8
	UINSEQ      APPkey = 1
)

var keyPrefix = map[APPkey]string{
	UINSEQ: "uinseq",
}

type APPkey int

// GenSeq 这里相当于将2^64分为8个区间，每次在区间递增一个值返回，暂时不考虑区间溢出的可能
func GenSeq(key APPkey) (uint64, error) {
	index := rand.New(Seed).Uint64() % shardingNum
	t := keyPrefix[key] + "_" + strconv.FormatUint(index, 10)
	seq, err := db.GenSeq(t)
	if err != nil {
		return 0, err
	}
	seq = index*(uint64)(math.Floor(math.MaxUint64/8)) + seq
	if seq < 1000 {
		seq += 1000
	}
	return seq, nil
}
