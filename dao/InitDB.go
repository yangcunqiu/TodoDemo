package dao

import (
	"TodoDemo/model/entity"
	"gorm.io/gorm"
	"log"
)
import "gorm.io/driver/mysql"

var db *gorm.DB

func InitDB() {
	// 初始化数据库连接
	dsn := "test:pwd1234@tcp(172.19.76.179)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	// 自动迁移
	err = db.AutoMigrate(entity.Todo{})
	if err != nil {
		log.Printf("gorm自动迁移失败, error%v", err)
	}
}
