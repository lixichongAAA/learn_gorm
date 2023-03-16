package main

import (
	"gorm.io/gorm"
)

type Info_ struct {
	gorm.Model
	Money int
	DogId uint
}

type Dog_ struct {
	gorm.Model
	Name     string
	Info_    Info_
	GirlGods []GirlGod_ `gorm:"many2many:dog_girl_god"`
}

type GirlGod_ struct {
	gorm.Model
	Name string
	Dogs []Dog_ `gorm:"many2many:dog_girl_god"`
}

func Many2Many() {
	//GLOBLE_DB.AutoMigrate(&Dog_{}, &GirlGod_{}, &Info_{})
	// i := Info_{
	// 	Money: 200,
	// }
	// g1 := GirlGod_{
	// 	Model: gorm.Model{
	// 		ID: 1,
	// 	},
	// 	Name: "女神一号",
	// }
	// g2 := GirlGod_{
	// 	Model: gorm.Model{
	// 		ID: 2,
	// 	},
	// 	Name: "女神二号",
	// }
	// d := Dog_{
	// 	Name:     "汪汪二号",
	// 	GirlGods: []GirlGod_{g1, g2},
	// 	Info_:    i,
	// } //现在创建的关系是dog中汪汪一号舔着女神一号和女神二号, 汪汪二号舔着女神一号和女神二号.代码更改执行来着.

	d := Dog_{
		Model: gorm.Model{
			ID: 3,
		},
	}
	//var girls []GirlGod_
	//GLOBLE_DB.Model(&d).Association("GirlGods").Find(&girls) //只查看d 添狗添的GirlGods列表不会打印d添狗
	//GLOBLE_DB.Model(&d).Preload("Dogs").Association("GirlGods").Find(&girls) //查看d添狗添的GirlGods列表，以及GirlGods列表中的GirlGod的Dogs列表
	//GLOBLE_DB.Model(&d).Preload("Dogs.Info_").Preload("Dogs").Association("GirlGods").Find(&girls)
	//GLOBLE_DB.Model(&d).Preload("Dogs.Info_").Association("GirlGods").Find(&girls)//等价于上面的一种形式,查看d添狗所舔GirlGods列表，以及GirlGods列表中GirlGod的Dogs的Info
	// GLOBLE_DB.Model(&d).Preload("Dogs", func(db *gorm.DB) *gorm.DB {
	// 	return db.Joins("Info_").Where("Money < ?", 1000)
	// }).Association("GirlGods").Find(&girls) //d号添狗所填的GirlGods列表中的GirGod中的Dogs中的Dog中的Info中的money < 10000 的

	//fmt.Println(girls)

	g1 := GirlGod_{
		Model: gorm.Model{
			ID: 1,
		},
	}

	g2 := GirlGod_{
		Model: gorm.Model{
			ID: 2,
		},
	}
	GLOBLE_DB.Model(&d).Association("GirlGods").Replace(&g1, &g2) //replace 替换已有关系为现在Replace中的参数
	//这句的意思是d号添狗都不舔以前所舔的女神了，现在就舔g1 和 g2
	//还有 Append, Delete, Clear 等,无需解释

}
