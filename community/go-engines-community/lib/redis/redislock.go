package redis

//go:generate mockgen -destination=../../mocks/lib/redis/redislock.go git.canopsis.net/canopsis/go-engines/lib/redis LockClient,Lock

import (
	"context"
	"github.com/bsm/redislock"
	"time"
)

type LockClient interface {
	Obtain(ctx context.Context, key string, ttl time.Duration, opt *redislock.Options) (Lock, error)
}

type Lock interface {
	Key() string
	Token() string
	TTL(ctx context.Context) (time.Duration, error)
	Refresh(ctx context.Context, ttl time.Duration, opt *redislock.Options) error
	Release(ctx context.Context) error
}

type lockClient struct {
	client *redislock.Client
}

func NewLockClient(redisClient redislock.RedisClient) LockClient {
	return &lockClient{
		client: redislock.New(redisClient),
	}
}

func (c *lockClient) Obtain(ctx context.Context, key string, ttl time.Duration, opt *redislock.Options) (Lock, error) {
	l, err := c.client.Obtain(ctx, key, ttl, opt)
	if err != nil {
		return nil, err
	}

	return l, nil
}
