package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() {
	var err error = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Dbdriver := os.Getenv("DB_DRIVER")
	if Dbdriver == "mysql" {
		// d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/belajar?parseTime=true")
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
		d, err := gorm.Open(Dbdriver, DBURL)
		if err != nil {
			panic(err)
		}
		db = d
	}
	if Dbdriver == "postgres" {
		// d, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=belajar sslmode=disable password=postgres")
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
		d, err := gorm.Open(Dbdriver, DBURL)
		if err != nil {
			panic(err)
		}
		db = d
	}
}

func GetDB() *gorm.DB {
	return db
}
