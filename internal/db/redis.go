package db

import (
    "github.com/go-redis/redis/v8"
    "context"
    "log"
)

var ctx = context.Background()

func ConnectRedis(addr string, password string, db int) *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password, // no password set
        DB:       db,       // use default DB
    })

    // 测试连接
    _, err := rdb.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("无法连接到Redis: %v", err)
    }

    return rdb
}