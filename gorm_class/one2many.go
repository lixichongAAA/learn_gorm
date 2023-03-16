package main

import "gorm.io/gorm"

type Dog struct {
	gorm.Model
	Name      string
	GirlGodID uint
}

type GirlGod struct {
	gorm.Model
	Name string
	Dogs []Dog
}

func one2many() {

}
