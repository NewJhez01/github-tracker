package repo

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func getRedisConn() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password
		DB:       0,  // use default DB
		Protocol: 2,
	})
}

func Set(ctx context.Context, report, key string) {
	rdb := getRedisConn()
	err := rdb.Set(ctx, key, report, 0).Err()
	if err != nil {
		panic(err)
	}
}

func Get(ctx context.Context, key string) {
	rdb := getRedisConn()
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println(val)
}
