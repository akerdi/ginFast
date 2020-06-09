package db

import (
	"ginFast/src/db/schema"
	"github.com/jinzhu/gorm"
)

func SetupTables(db *gorm.DB) error {
	if !db.HasTable(&schema.User{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&schema.User{}).Error; err != nil {
			return err
		}
	}
	return  nil
}