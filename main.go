package main

import (
	"crowdfunding/user"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	input := "2022-04-12 15:07:00"
	layout := "2006-01-02 15:04:05"
	date, _ := time.Parse(layout, input)
	user := user.User{
		ID:         0,
		Name:       "Muhammad",
		Occupation: "Data Scientist",
		Email:      "kumbara@gmail.com",
		Password:   "test123",
		Avatar:     "avatar.jpg",
		Role:       "user",
		CreatedAt:  date,
		UpdatedAt:  date,
	}

	userRepository.Save(user)

}
