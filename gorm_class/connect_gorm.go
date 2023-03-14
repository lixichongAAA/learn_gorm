package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:123456@tcp(127.0.0.1:3306)/gorm_class?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize: 171,                                                                                   // string 类型字段的默认长度
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "gva_",
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	type User struct {
		Name string
	}

	M := db.Migrator()
	if M.HasTable(&User{}) {
		M.DropTable(&User{})
	} else {
		M.CreateTable(&User{})
	}
}
