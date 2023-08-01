package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/redis/go-redis/v9"
	"log"
)

func Init(datasource string) *gorm.DB {
	db, err := gorm.Open("mysql", datasource) //打开mysql数据库连接
	if err != nil {
		log.Println("Gorm创建MySQL连接时Error：", err)
	}
	return db
}

func InitRedis(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
}
