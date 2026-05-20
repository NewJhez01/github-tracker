package repo

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

type Repo struct {
	client *redis.Client
}

func NewCacheRepo(c *redis.Client) Repo {
	return Repo{
		client: c,
	}
}

func (r *Repo) Set(ctx context.Context, report, key string) error {
	err := r.client.Set(ctx, key, report, 0).Err()
	if err != nil {
		return errors.New("failed to set key in redis prev: " + err.Error())
	}
	return nil
}

func (r *Repo) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", errors.New("failed to get redis val with key: " + key + " prev: " + err.Error())
	}

	return val, nil
}
