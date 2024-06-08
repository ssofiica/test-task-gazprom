package connection

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ssofiica/test-task-gazprom/config"
)

func InitRedis(cfg *config.Project, dbNum int) *redis.Client {
	redisConfig := cfg.Redis

	redisAddress := fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)
	r := redis.NewClient(&redis.Options{
		DB:   dbNum,
		Addr: redisAddress,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis", err.Error())
	}

	fmt.Println("Redis connected successfully")
	return r
}
