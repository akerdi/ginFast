package db

import (
	"ginFast/src/db/schema"
	"github.com/jinzhu/gorm"
	"github.com/shaohung001/ginFastApp"
)

func SetupTables(db *gorm.DB) error {
	if !db.HasTable(&schema.User{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&schema.User{}).Error; err != nil {
			return err
		}
	}
	return  nil
}

var RedisConnect *ginFastApp.RedisClient

func SetupRedis(redisConnect *ginFastApp.RedisClient)  {
	RedisConnect = redisConnect
}

func RedisGetRangeByKey(key string, left, right int64) ([]string, error) {
	if right == 0 {
		right = 10
	}
	results, err := RedisConnect.GetClient().LRange(key, left, right).Result()
	return results, err
}