package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started succcessfully: pong message = {%s}\n", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveUrlMapping(shortUrl string, originUrl string, userId string) {
	err := storeService.redisClient.Set(shortUrl, originUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originUrl: %s", err, shortUrl, originUrl))
	}
}

func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed retrieving key url | Error: %v - shortUrl: %s", err, shortUrl))
	}
	return result
}
