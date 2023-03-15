package main

import "fmt"

func TestFind() {
	var result map[string]interface{}
	GLOBLE_DB.Model(&TestUser{}).First(&result)
	fmt.Println(result)
}
