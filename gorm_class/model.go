package main

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	UUID uint      `gorm:"primaryKey"`
	Time time.Time `gorm:"column:my_time"`
}

type TestUser struct {
	gorm.Model
	Name string `gorm:"default:lxc"`
	Age  uint8  `gorm:"comment:age111"`
}

func TestUserCreate() {
	GLOBLE_DB.AutoMigrate(&TestUser{})
}
