package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Database Connection Faliure\n%s", err.Error()))
	}

	return db
}
