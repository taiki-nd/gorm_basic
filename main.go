package main

import (
	"fmt"
	"gorm_basic/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id        int64
	FirstName string
	LastName  string
	Email     string
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Config.DbUserDevelop, config.Config.DbPasswordDevelop, config.Config.DbHostDevelop,
		config.Config.DbPortDevelop, config.Config.DbNameDevelop)

	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	//
	// create record
	//
	/*
		user := User{
			FirstName: "Taiki",
			LastName:  "Noda",
			Email:     "hoge@hogehoge",
		}
		db.Create(&user)
	*/

	//
	// update user
	//
	user := User{
		Id:        1,
		FirstName: "Taiki updated",
		LastName:  "Noda updated",
		Email:     "hoge@hogehoge updated",
	}
	db.Updates(&user)

}
