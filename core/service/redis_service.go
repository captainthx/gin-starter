package service

import (
	"gin-starter/config"
	"gin-starter/logs"
	"time"

	"github.com/redis/go-redis/v9"
)

// function set ค่าใน redis server
func SetValueInRedis(redisKey string, value string, duration time.Duration) error {
	err := config.RedisClient.Set(config.Ctx, redisKey, value, duration).Err()
	if err != nil {
		logs.Warn("SetValueInRedis-[block].(Error setting value in redis). error: " + err.Error())
		return err
	}
	return nil
}

// funcion get ค่าใน redis server (หากไม่เกิด error จะคืนค่า string และ error เป็น nil)
func GetValueInRedis(redisKey string) (string, error) {

	val, err := config.RedisClient.Get(config.Ctx, redisKey).Result()

	if err != redis.Nil {
		logs.Warn("GetValueInRedis-[block].(Key not found in redis). error: " + redisKey)
		return "", nil
	} else if err != nil {
		logs.Warn("GetValueInRedis-[block].(Can not get value form redis). error: " + err.Error())
		return "", nil
	}
	return val, nil
}
