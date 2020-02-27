package common

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Cfg   *viper.Viper
	DB    *gorm.DB
	Redis *redis.Client
	Email *yuEmail
	L     *logrus.Entry
)
