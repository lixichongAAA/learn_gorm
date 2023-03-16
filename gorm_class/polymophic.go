package main

// 多态适用于 has one 和 has many 的关系
// type Jiazi struct {
// 	ID          uint
// 	Name        string
// 	XiaoFengChe []XiaoFengChe `gorm:"polymorphic:Owner;polymorphicValue:huhu"` //polymorphicValue 用于改变拥有者类型的名称,比如原来小风车拥有者的类型为Jiazi 现在为huhu
// } //has many

// type Yujie struct {
// 	ID          uint
// 	Name        string
// 	XiaoFengChe XiaoFengChe `gorm:"polymorphic:Owner;"`
// } //has a

// type XiaoFengChe struct {
// 	ID        uint
// 	Name      string
// 	OwnerType string //OwnerType顾名思义就是这个xiaofengche拥有者的type,
// 	OwnerID   uint   //OwnerID 就可以在确定拥有者的类型type情况下,根据OwnerID确认拥有者了
// }

// 连接表相关
type JiaZi struct {
	ID          uint
	Name        string
	XiaoFengChe []XiaoFengChe `gorm:"many2many:jiazi_fengche;foreignKey:Name;references:FCName"`
	//连接表的两个列名分别为jia_zi_name, xiao_feng_che_fc_name 分别列的是Jiazi结构体的Name和xiaofengche结构体的FCName
	//joinForeignKey joinReferences指定连接表的列的名字, 比如`gorm:"many2many:jiazi_fengche;foreignKey:Name;joinForeignKey: aaa; references:FCName;joinReferences: bbb"`
	//就是更改了对应的列名为aaa bbb;
}

type XiaoFengChe struct {
	ID     uint
	FCName string
}

// has many关系指定外键与引用,这样可以少个链子
// 一个jiazi有many 小风车 正常来说，xiaofengche内应该有一个指向jiazi的字段，现在就是通过指定外键和引用从而也达到了指定的目的
// type JiaZi struct {
// 	ID          uint
// 	Name        string
// 	XiaoFengChe []XiaoFengChe `gorm:"foreignKey:JiaZiName;references:Name"`
// }

// type XiaoFengChe struct {
// 	ID        uint
// 	JiaZiName string
// }

func Polymophic() {
	//GLOBLE_DB.AutoMigrate(&Jiazi{}, &Yujie{}, &XiaoFengChe{})
	// GLOBLE_DB.Create(&Jiazi{Name: "夹子2", XiaoFengChe: []XiaoFengChe{
	// 	{Name: "小风车"},
	// 	{Name: "小风车2"},
	// }})
	// GLOBLE_DB.Create(&Yujie{Name: "御姐2", XiaoFengChe: XiaoFengChe{
	// 	Name: "大风车",
	// }})

	//连接表
	GLOBLE_DB.AutoMigrate(&JiaZi{}, &XiaoFengChe{})
	GLOBLE_DB.Create(&JiaZi{Name: "小夹子", XiaoFengChe: []XiaoFengChe{
		{FCName: "小风车1"},
		{FCName: "小风车2"},
		{FCName: "小风车3"},
	}})
}
