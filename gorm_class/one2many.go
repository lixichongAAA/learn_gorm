package main

import (
	"fmt"

	"gorm.io/gorm"
)

type Info struct {
	gorm.Model
	Money int
	DogID uint
}

type Dog struct {
	gorm.Model
	Name      string
	GirlGodID uint
	Info      Info
}

type GirlGod struct {
	gorm.Model
	Name string
	Dogs []Dog
}

func one2many() {
	//GLOBLE_DB.AutoMigrate(&Dog{}, &GirlGod{})
	// d1 := Dog{Name: "汪汪一号"}
	// d2 := Dog{Name: "汪汪二号"}
	// g := GirlGod{
	// 	Name: "女神一号",
	// 	Dogs: []Dog{d1, d2},
	// }

	// GLOBLE_DB.Create(&g)

	//var girl GirlGod
	// GLOBLE_DB.First(&girl)
	// fmt.Println(girl) //这个打印{{1 2023-03-16 11:22:36.855 +0800 CST 2023-03-16 11:22:36.855 +0800 CST {0001-01-01 00:00:00 +0000 UTC false}} 女神一号 []}
	//在girl的dogs列表中没有查出来dogs

	//预加载
	// GLOBLE_DB.Preload("Dogs").First(&girl) //此时通过preload打印出了dogs dogs是以Dog的结构体的形式打印出来的
	// fmt.Println(girl)

	//带条件的预加载
	// GLOBLE_DB.Preload("Dogs", "name = ?", "汪汪一号").Find(&girl)
	// fmt.Println(girl)

	//自定义SQL预加载
	// GLOBLE_DB.Preload("Dogs", func(db *gorm.DB) *gorm.DB {
	// 	return db.Where("name = ?", "汪汪一号")
	// }).First(&girl)
	// fmt.Println(girl)

	//链式预加载，增加了Info结构体
	// GLOBLE_DB.AutoMigrate(&Info{})
	// GLOBLE_DB.Create(&Info{
	// 	Money: 100,
	// 	DogID: 1,
	// })
	// GLOBLE_DB.Create(&Info{
	// 	Money: 20000,
	// 	DogID: 2,
	// })

	//总的来说链式预加载,只能在preload的这层查询这层的条件
	var girl GirlGod
	//GLOBLE_DB.Preload("Dogs").First(&girl)//此时的预加载只是加载出了dogs结构体
	//GLOBLE_DB.Preload("Dogs.Info").Preload("Dogs").First(&girl) // 等同于 GLOBLE.DB.Preload("Dogs.Info").First(&girl)推荐用前一种形式，后者无法用Dogs结构体的条件查询
	/*111*/ //GLOBLE_DB.Preload("Dogs.Info", "money > ?", 200).Preload("Dogs", "name = ?", "汪汪一号").First(&girl)

	GLOBLE_DB.Preload("Dogs", func(db *gorm.DB) *gorm.DB {
		return db.Joins("Info").Where("money > ?", 200)
	}).First(&girl) //这种形式只会打印汪汪二号,和111形式结果有所区别,111形式在第二个preload中不加where条件时也会查出汪汪一号，不过因为汪汪一号不满足条件，其部分字段显示为零

	fmt.Println(girl)
}
