package main

import (
	"bluebell_sly/dao/postgres"
	"bluebell_sly/dao/redis"
	"bluebell_sly/logger"
	"bluebell_sly/settings"
	"fmt"
)

func main() {
	// 初始化配置
	if err := settings.Init(); err != nil {
		fmt.Println("settings初始化 failed,err:", err)
		return
	}
	logger.Init(settings.Conf.LogConfig, settings.Conf.Mode)
	postgres.Init(settings.Conf.PostgresqlConfig)
	defer postgres.Close()
	err := redis.Init(settings.Conf.RedisConfig)
	if err != nil {
		fmt.Println("redis connect failed, err:", err)
	}
	defer redis.Close()

}
