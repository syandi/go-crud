package config

import (
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	// d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/belajar?parseTime=true")
	d, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=belajar sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
