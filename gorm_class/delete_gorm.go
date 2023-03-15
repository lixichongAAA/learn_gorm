package main

func TestDelete() {
	var users []TestUser
	GLOBLE_DB.Where("Name = ?", "Lxc").Delete(&users)            //当存在deleted_at 时为软删
	GLOBLE_DB.Unscoped().Where("Name = ?", "Lxc").Delete(&users) //硬删
}
