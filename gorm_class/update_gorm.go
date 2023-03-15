package main

func TestUpdate() {
	//update 只更新你选择的字段
	//updates 更新所有字段,有两种形式, 一种为map， 一种为结构体，结构体零值不参与更新
	//save 更新所有内容，包括零值
	var Users []TestUser

	// dbRes := GLOBLE_DB.Where("name LIKE ?", "%L%").Find(&Users)
	// for k := range Users {
	// 	Users[k].Age = 18
	// }

	// dbRes.Save(&Users)

	GLOBLE_DB.Find(&Users).Updates(map[string]interface{}{
		"Name": "",
		"Age":  12,
	})

	var users TestUser
	GLOBLE_DB.Find(&users).Updates(TestUser{
		Name: "Lxc",
		Age:  0,
	})

}
