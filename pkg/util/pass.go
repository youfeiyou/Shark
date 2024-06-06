package util

import (
	"crypto/md5"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"time"
)

const (
	letterbytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

var seed = rand.NewSource(time.Now().UnixNano())

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
		b[i] = letterbytes[rand.New(seed).Intn(len(letterbytes))]
	}
	return b
}
