package utils

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/iamaul/fatbellies/config"
	"github.com/sirupsen/logrus"
)

var Redis *redis.Client

func ConnectRedis(c *config.Configuration) (*redis.Client, error) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort),
		Password: c.RedisPassword,
		DB:       0,
	})

	_, err := Redis.Ping().Result()
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info("Connected with Redis.")

	return Redis, err
}
