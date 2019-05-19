package data

import (
	"encoding/json"
	"mess-app/internal/core"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type redisCache struct {
	c *redis.Client
}

func (r *redisCache) Get(key string, val interface{}) error {
	result, err := r.c.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			return core.ErrResourceNotFound
		}
		return errors.Wrap(err, "Failed to get value from redis")
	}
	if err = json.Unmarshal([]byte(result), &val); err != nil {
		return errors.Wrap(err, "Failed to unmarshal")
	}
	return nil
}
func (r *redisCache) Set(key string, val interface{}, d time.Duration) error {
	data, err := json.Marshal(val)
	if err != nil {
		return errors.Wrap(err, "Failed to marshal")
	}
	err = r.c.Set(key, data, d).Err()
	if err != nil {
		return errors.Wrap(err, "Failed to set value to redis")
	}
	return nil
}

func (r *redisCache) IfExists(key string) (bool, error) {
	val, err := r.c.Exists(key).Result()
	return val == 1, err
}
