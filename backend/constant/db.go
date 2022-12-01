package constant

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"hospitie/core"
)

var lock = &sync.Mutex{}

var DB *gorm.DB

func DBInstance() *gorm.DB {
	if DB == nil {
		print("create DB")
		lock.Lock()
		defer lock.Unlock()
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PATH"), os.Getenv("DB_NAME"))
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			core.ErrorHandler(err)
		} else {
			DB = db
		}
	}
	return DB
}
