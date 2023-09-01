package cache

import (
	"context"
	"go-init/global"
	"time"
)

type Cache struct{}

// GetKV 根据获取值
func (c *Cache) GetKV(ctx context.Context, key string) (string, error) {
	result, err := global.CaChe.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

// SetKV 设置键值存入 redis
func (c *Cache) SetKV(ctx context.Context, key string, data []byte, expire time.Duration) error {
	_, err := global.CaChe.Set(ctx, key, data, expire).Result()
	if err != nil {
		return err
	}
	return nil
}

// DeleteKV 删除键
func (c *Cache) DeleteKV(ctx context.Context, keyPattern string) error {
	// Find keys matching the pattern
	keys, err := global.CaChe.Keys(ctx, keyPattern).Result()
	if err != nil {
		return err
	}

	// Ensure there are keys to delete
	if len(keys) == 0 {
		return nil
	}

	// Delete the keys
	_, err = global.CaChe.Del(ctx, keys...).Result()
	if err != nil {
		return err
	}

	return nil
}
