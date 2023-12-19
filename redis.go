package handlers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func init() {

	conn := fmt.Sprintf("%s:%s", os.Getenv("REDIS_SERVICE_HOST"), os.Getenv("REDIS_SERVICE_PORT")) // KUBERNETES

	Client = redis.NewClient(&redis.Options{
		Addr:     conn,
		Password: "",
		DB:       0,
	})
}

func SetRedisWithTime(key string, value []byte, timeAccess bool, dr int64) error {
	duration := time.Duration(dr) * time.Minute
	if timeAccess {

		rduration := os.Getenv("RedisDuration")

		n, _ := strconv.Atoi(rduration)
		duration = time.Duration(n) * time.Minute
	}
	err := Client.Set(key, value, duration).Err()
	if err != nil {
		return err
	}
	return err
}

func SetRedis(key string, value []byte) error {
	err := Client.Set(key, value, 0).Err()
	if err != nil {
		fmt.Println("ERROR CREATE REDIS : ", err)
	}
	return err
}

func GetRedis(key string) (string, error) {
	val, err := Client.Get(key).Result()
	return val, err
}

func DelRedis(key string) {
	Client.Del(key)
}

func DelRedisWithLoop(key string) {
	val := Client.Keys(key).Val()
	for _, s := range val {
		fmt.Println(Client.Del(s))
	}

}
