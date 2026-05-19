package repo

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Repo struct {
	client *redis.Client
}

func NewCacheRepo(c *redis.Client) *Repo {
	return &Repo{
		client: c,
	}
}

func (r *Repo) Set(ctx context.Context, report, key string) {
	err := r.client.Set(ctx, key, report, 0).Err()
	if err != nil {
		fmt.Println("failed to set the value in redis")
	}
}

func (r *Repo) Get(ctx context.Context, key string) string {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err.Error())
	}

	return val
}
