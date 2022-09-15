package main

import (
	"GoApi/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	dsn := "host=localhost user=postgres dbname=gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err = db.AutoMigrate(&User{}); err != nil {
		log.Fatal(err)
	}

	server.Handle("/", func(user *User) *User {
		db.Save(user)
		return user
	})
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
