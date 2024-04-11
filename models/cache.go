package models

import (
	"fmt"

	"tiny-sec-kill/utils"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func InitCache() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", utils.RdsIp, utils.RdsPort),
		Password: utils.RdsPassword, // no password set
		DB:       0,                 // use default DB
	})
	fmt.Println(rdb.Ping().Result())
}
