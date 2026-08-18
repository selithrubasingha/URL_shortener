package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)


// Define the struct wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
}

// Top level declarations for the storeService and Redis context
var (
	storeService  = &StorageService{}

	//ctx is a blank, default context that Redis requires you to pass in to make sure requests don't hang forever
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

	pong, err := redisClient.Ping(ctx).Result() // ping checks if the database if awake .

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	fmt.Println("Redis connection successful:", pong)

	storeService.redisClient = redisClient
	return storeService

}

/* We want to be able to save the mapping between the originalUrl 
and the generated shortUrl url
*/

func SaveUrlMapping(shortUrl string , originalUrl string , userId string)  {

	// Set(ctx , key , value , expiration time) value - the actual heavy data you want to store  .
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))

	}

}

/*
We should be able to retrieve the initial long URL once the short 
is provided. This is when users will be calling the shortlink in the 
url, so what we need to do here is to retrieve the long url and
think about redirect.
*/

func RetrieveInitialUrl(shortUrl string) string {

	// Get(ctx , key) - key is the shortUrl that we generated and stored in the cache
	result , err := storeService.redisClient.Get(ctx, shortUrl).Result()

	 if err != nil {
		panic(fmt.Sprintf("Failed retrieving key url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}

	return result

}

