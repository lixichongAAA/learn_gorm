package main

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GLOBLE_DB *gorm.DB

func main() {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:123456@tcp(127.0.0.1:3306)/gorm_class?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize: 171,                                                                                   // string 类型字段的默认长度
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	GLOBLE_DB = db

	//TestUserCreate()
	//CreatedTest()
	//TestFind()
	//TestUpdate()
	//TestDelete()
	//TestRaw()
	//one2one()
	//one2many()
	//Many2Many()
	//TestTranscaction()
	//Custome()
	Polymophic()
}
