package config

import (
	"github.com/kangyueyue/short-url/internal/infrastructure/store"
	"github.com/redis/go-redis/v9"
)

// ProvideStore wire
func ProvideStore(cfg *Config) (*store.Store, error) {
	return store.NewStore(cfg.DbConfig)
}

// ProvideRedisClient wire
func ProvideRedisClient(cfg *Config) *redis.Client {
	return redis.NewClient(cfg.RedisConfig)
}
