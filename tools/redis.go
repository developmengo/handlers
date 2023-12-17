package tmp

import (
	"fmt"
	"os"

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

func SetRedis(key string, value []byte) error {
	err := Client.Set(key, value, 0).Err()
	if err != nil {
		fmt.Println("ERROR CREATE OPERATOR REDIS : ", err)
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
