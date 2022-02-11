package dao

import (
	"Simple-login-demo/setting"
	"github.com/go-redis/redis"
)

var RedisDB *redis.Client

func InitRedis(cfg *setting.RedisConfig) (err error) {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: "",     // no password set
		DB:       cfg.DB, // use default DB
	})
	_, err = RedisDB.Ping().Result()
	return
}

func RedisClose() {
	RedisDB.Close()
}
