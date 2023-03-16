package main

// belongs to
// type Dog struct {
// 	gorm.Model
// 	Name      string
// 	GirlGodID uint
// 	GirlGod   GirlGod
// }

// type GirlGod struct {
// 	gorm.Model
// 	Name string
// }

// has one
// type Dog struct {
// 	gorm.Model
// 	Name      string
// 	GirlGodID uint
// }

// type GirlGod struct {
// 	gorm.Model
// 	Name string
// 	Dog  Dog
// }

func one2one() {
	//belongs to
	// g := GirlGod{
	// 	Model: gorm.Model{
	// 		ID: 1,
	// 	},
	// 	Name: "ns",
	// }

	// d := Dog{
	// 	Model: gorm.Model{
	// 		ID: 1,
	// 	},
	// 	Name:    "dd",
	// 	GirlGod: g,
	// }
	// GLOBLE_DB.Create(&d)//创建dogs时也会创建girlgod,先有girlgod后有dog
	//GLOBLE_DB.AutoMigrate(&Dog{}) //belongs to 在创建dogs时也会创建GirlGod,不同于结构体嵌套，dogs内没有girlgod的字段，而是会创建girlgod

	// has a
	// d := Dog{
	// 	Model: gorm.Model{
	// 		ID: 2,
	// 	},
	// 	Name: "dd",
	// }

	// g := GirlGod{
	// 	Model: gorm.Model{
	// 		ID: 2,
	// 	},
	// 	Name: "ns",
	// 	Dog:  d,
	// }
	// GLOBLE_DB.Create(&g) //注意，尽管grilgod内has a dog但是girlgod是不依赖dog而存在的，即girlgod内没有dog的字段
	// GLOBLE_DB.AutoMigrate(&GirlGod{}, &Dog{})
	// var girl GirlGod
	// GLOBLE_DB.Preload("Dog").First(&girl, 1) // 预加载，查询girlgod时会带出她所拥有的dog. 在查dog时同样也可以预加载girlgod
	// fmt.Println(girl)
}
