package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func DBConnect() *gorm.DB{
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if nil != err {
		log.Printf("ERROR: failed dbConnect open postgres, message: %s", err.Error())
		log.Fatal(err)
	}

	return db
}