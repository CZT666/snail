package dao

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"student_bakcend/settings"
	"time"
)

var RedisDB *redis.Client

func InitRedis(config *settings.RedisConfig) (err error) {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: "",
		DB:       config.DB,
		PoolSize: config.PoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = RedisDB.Ping(ctx).Result()
	return
}

func CloseRedis() {
	err := RedisDB.Close()
	if err != nil {
		log.Printf("Close redis error: %v\n", err)
	}
	return
}
