package cache

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/lr2021/recruit-backend/general/config"
	"log"
	"sync"
)

var (
	rdb *redis.Client
	inited bool
	m      sync.RWMutex
)

func Init() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[Init] redis has inited")
		log.Println(err.Error())
		return
	}

	rdb = redis.NewClient(&redis.Options{
		Addr: config.REDIS_ADDR,
		Password: config.REDIS_PASSWORD,
		DB: 0,
		PoolSize: 100,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		panic("[init] Error when init redis: " + err.Error())
	}
}

func HasInit() bool {
	if inited {
		return true
	}

	return false
}

func GetRDB() *redis.Client {
	if !inited {
		panic("[init] Redis does not init")
	}

	return rdb
}
