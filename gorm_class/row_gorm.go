package main

import "fmt"

type UserInfo2 struct {
	Name string
	Age  uint8
}

func TestRaw() {
	var users []UserInfo2
	GLOBLE_DB.Raw("SELECT * FROM test_users WHERE name = ?", "Lxc").Scan(&users)
	fmt.Println(users)
}
