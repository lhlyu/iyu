package module

import (
	"github.com/go-redis/redis"
	"github.com/lhlyu/iyu/common"
	"log"
	"time"
)

type rds struct {
}

func (rds) SetUp() {
	c := &redisConf{}
	if err := common.Cfg.UnmarshalKey("redis", c); err != nil {
		log.Fatal("db setup is err:", err)
	}
	setRedis(c)
}

// redis模块
var RedisModule = rds{}

type redisConf struct {
	Addr        string `json:"addr"`
	Password    string `json:"password"`
	Database    int    `json:"database"`
	IdleTimeout int    `json:"idleTimeout"`
}

func setRedis(r *redisConf) {

	client := redis.NewClient(&redis.Options{
		Addr:        r.Addr,
		Password:    r.Password,
		DB:          r.Database,
		IdleTimeout: time.Duration(r.IdleTimeout) * time.Second,
	})
	if _, err := client.Ping().Result(); err != nil {
		log.Fatal("redis connect is fail,err:", err)
	}
	common.Redis = client
}
