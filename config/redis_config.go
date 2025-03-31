package config

import (
	"context"
	"gin-starter/logs"

	"github.com/redis/go-redis/v9"
)

// สร้างตัวแปร global สำหรับเก็บ Redis client
var RedisClient *redis.Client

// สร้าง context.Background() ที่ไม่มีการหมดเวลา
var Ctx = context.Background()

// ฟังก์ชั่นการตั้งค่า Redis client
func InitRedisClient(redisServer string) {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisServer,
		Password: "",
		DB:       0,
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		logs.Error("InitRedisClient-[block].(Can not connect to redis). error: " + err.Error())
	}
	logs.Info("InitRedisClient-[next].(Redis connected successfully✅)")
}
