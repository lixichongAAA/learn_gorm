package main

import (
	"fmt"
)

type UserInfo struct {
	Name string
	Age  uint8
}

func TestFind() {
	var User = []UserInfo{}
	GLOBLE_DB.Model(&TestUser{}).Where("name <> ?", "L").Find(&User)
	fmt.Println(User)
}
