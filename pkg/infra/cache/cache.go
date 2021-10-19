package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"note-manager/pkg/infra/config"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	once sync.Once
	rdb  *redis.Client
	ctx  context.Context
)

// Cache is handler of cache
type Cache struct {
	key string
}

// NewCache new a cache
func NewCache() *Cache {
	once.Do(func() {
		address := fmt.Sprintf("%v:%v", config.GetRdbAdress(), config.GetRdbPort())
		rdb = redis.NewClient(&redis.Options{
			Addr:     address,
			Password: config.GetRdbPassword(),
			DB:       0, // use default DB
		})
		ctx = context.Background()
	})
	c := &Cache{}
	return c
}

// Set value
func (c *Cache) Set(val interface{}) error {
	key := c.key
	data, _ := json.Marshal(val)
	if err := rdb.Set(ctx, key, data, time.Hour).Err(); err != nil {
		return err
	}
	return nil
}

func (c *Cache) get(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

// Key store key
func (c *Cache) Key(format string, a ...interface{}) *Cache {
	h := &Cache{}
	h.key = fmt.Sprintf(format, a...)
	return h
}

// Get get value
func (c *Cache) Get(v interface{}) error {
	val, err := c.get(c.key)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), &v)
	return err
}
