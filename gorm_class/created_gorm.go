package main

import "fmt"

func CreatedTest() {
	dbrsp := GLOBLE_DB.Omit("Name").Create(&[]TestUser{
		{Name: "Lxc", Age: 18},
		{Name: "sss", Age: 122},
		{Name: "woshi", Age: 23},
	})
	fmt.Println(dbrsp.Error, dbrsp.RowsAffected)
}
