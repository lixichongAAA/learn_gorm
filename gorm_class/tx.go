package main

type TMG struct {
	ID   uint
	Name string
}

func TestTranscaction() {
	// GLOBLE_DB.AutoMigrate(&TMG{})
	// flag := true
	// GLOBLE_DB.Transaction(func(tx *gorm.DB) error {
	// 	tx.Create(&TMG{Name: "汉子"})
	// 	tx.Create(&TMG{Name: "威武雄壮"})
	// 	tx.Create(&TMG{Name: "一只鸡鸡"})
	// 	if flag {
	// 		return nil
	// 	} else {
	// 		return errors.New("躲汉子") //这会导致事务回滚,即不会创建表中的3条记录 error不为nil
	// 	}
	// })

	// GLOBLE_DB.AutoMigrate(&TMG{})
	// GLOBLE_DB.Transaction(func(tx *gorm.DB) error {
	// 	tx.Create(&TMG{Name: "威武雄壮"})
	// 	tx.Create(&TMG{Name: "一只雄鹰"})
	// 	tx.Transaction(func(tx *gorm.DB) error {
	// 		tx.Create(&TMG{Name: "躲汉子"})
	// 		return errors.New("躲汉子") //因为返回了错误,所以该事务内的创建语句不会被创建
	// 	})
	// 	return nil
	// })

	// //手动事务,在手动事务中，你要做的事必须在commit之前，commit 之后的语句不会被执行,且会编译报错
	// GLOBLE_DB.AutoMigrate(&TMG{})
	// flag := false
	// GLOBLE_DB.Create(&TMG{Name: "000"}) // 该语句会被执行
	// tx := GLOBLE_DB.Begin()             //开启一个手动事务
	// tx.Create(&TMG{Name: "aaa"})
	// tx.Create(&TMG{Name: "bbb"})
	// tx.Create(&TMG{Name: "ccc"})
	// if flag {
	// 	tx.Rollback() //回滚
	// }
	// tx.Commit()
	// tx.Create(&TMG{Name: "ddd"}) //该语句不会被执行,且会报错

	GLOBLE_DB.AutoMigrate(&TMG{})
	tx := GLOBLE_DB.Begin()
	tx.Create(&TMG{Name: "111"})
	tx.Create(&TMG{Name: "222"})

	tx.SavePoint("hahaha")
	tx.Create(&TMG{Name: "被丢弃"})
	tx.RollbackTo("hahaha")

	tx.Commit()
}
