package util

import (
	"crypto/md5"
	"github.com/google/uuid"
	"hash/fnv"
	"log"
	"math/rand"
	"time"
)

const (
	letterbytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

var Seed = rand.NewSource(time.Now().UnixNano())

func UUID() ([]byte, error) {
	id, err := uuid.New().MarshalBinary()
	if err != nil {
		log.Fatalf("gen uuid error: %v", err)
		return nil, err
	}
	return id, nil
}

func Md5(val []byte) []byte {
	h := md5.New()
	h.Write(val)
	return h.Sum(nil)
}

func RandString(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterbytes[rand.New(Seed).Intn(len(letterbytes))]
	}
	return b
}

func HashString(s string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}
