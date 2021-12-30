package config

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var RedisCtx = context.Background()
var Rdb = redis.NewClient(&redis.Options{
	Addr:     "192.168.0.32:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

var Secret = []byte("secret")
