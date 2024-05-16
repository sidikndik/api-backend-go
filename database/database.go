package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func DBConnect() *gorm.DB {
// 	dsn := "host=localhost user=postgres password=penc3r4han dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if nil != err {
// 		log.Printf("ERROR: failed dbConnect open postgres, message: %s", err.Error())
// 		log.Fatal(err)
// 	}

// 	return db
// }

func DBConnect() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/golangdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if nil != err {
		log.Printf("ERROR: failed dbConnect open mysql, message: %s", err.Error())
		log.Fatal(err)
	}

	return db
}
