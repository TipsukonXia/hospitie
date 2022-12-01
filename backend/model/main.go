package model

import (
	"hospitie/constant"
)

func Migrate() error {
	db := constant.DBInstance()
	err := db.AutoMigrate(&User{}, &Hospitie{}, &Car{}, &History{})
	return err
}
