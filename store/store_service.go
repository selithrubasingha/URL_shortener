package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)


// Define the struct wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
}

// Top level declarations for the storeService and Redis context
var (
	storeService  = &StorageService{}
	ctx = context.Background()
)


// Note that in a real world usage, the cache duration shouldn't have  
// an expiration time, an LRU policy config should be set where the 
// values that are retrieved less often are purged automatically from 
// the cache and stored back in RDBMS whenever the cache is full

const CacheDuration = 6 * time.Hour


func InitializeStore() *StorageService {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := redisClient.Ping(ctx).Result()

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	fmt.Println("Redis connection successful:", pong)

	storeService.redisClient = redisClient
	return storeService

}

