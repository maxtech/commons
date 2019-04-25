package commons

import (
    "github.com/go-redis/redis"
    "log"
)

func InitRedisClient(address string, password string, dbNum int) *redis.Client {
    redisClient := redis.NewClient(&redis.Options{
        Addr:     address,
        Password: password, // no password set
        DB:       dbNum,    // use default DB
    })
    return redisClient
}

func TestRedisClient(redisClient *redis.Client) bool {
    _, err := redisClient.Ping().Result()
    if err != nil {
        log.Println("redis connect error: ", err)
        return false
    }
    return true
}
