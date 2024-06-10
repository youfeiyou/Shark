package db

import (
	"github.com/go-redis/redis"
	"reflect"
	"testing"
)

var (
	addr = "39.108.64.37:6379"
)

func TestNewRedisClient(t *testing.T) {
	cli := redis.Client{}
	tests := []struct {
		name string
		want *redisClinet
	}{
		{
			"rdis client",
			&redisClinet{
				&cli,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRedisClient(addr); !reflect.DeepEqual(got.cli.Options().Addr, addr) {
				t.Errorf("NewRedisClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisClinet_Hmget(t *testing.T) {
	type fields struct {
		cli *redis.Client
	}
	type args struct {
		key    string
		fields []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "test hmset",
			fields: fields{
				cli: NewRedisClient(addr).cli,
			},
			args: args{
				key:    "test_key",
				fields: []string{"1", "2"},
			},
			want: map[string]string{
				"1": "123",
				"2": "456",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewRedisClient(addr)
			kv := map[string]interface{}{
				"1": []byte("123"),
				"2": []byte("456"),
			}
			c.HMSet("test_key", kv)
			got, err := c.HMGet(tt.args.key, tt.args.fields...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hmget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got: %v, %v %v", len(got), got["1"], got["2"])
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hmget() got = %v, want %v", got, tt.want)
			}
		})
	}
}
