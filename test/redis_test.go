package test

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

var ctx = context.Background() //上下文
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // 没有密码，默认值
	DB:       0,  // 默认DB 0
})

func TestRedis(t *testing.T) {
	err := rdb.Set(ctx, "k1", "v2", time.Second*10).Err()
	if err != nil {
		t.Fatal(err)
	}
}
