package cache

import (
	. "github.com/hotrungnhan/go-fiber-template/pkg/configs"
	"github.com/hotrungnhan/go-fiber-template/pkg/utils"

	"github.com/go-redis/redis/v8"
)

// RedisConnection func for connect to Redis server.
func RedisConnection() (*redis.Client, error) {

	// Build Redis connection URL.
	redisConnURL, err := utils.ConnectionURLBuilder("redis")
	if err != nil {
		return nil, err
	}

	// Set Redis options.
	options := &redis.Options{
		Addr:     redisConnURL,
		Password: Get().CACHE.PASSWORD,
		DB:       Get().CACHE.DB_NUMBER,
	}

	return redis.NewClient(options), nil
}
