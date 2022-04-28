package main

import (
	"database/sql"
	"fmt"
	"gorm_basic/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName sql.NullString `gorm:"type:VARCHAR(30); default:'Taro'"`
	LastName  sql.NullString `gorm:"size:100; null"`
	Email     sql.NullString `gorm:"unique; not null"`
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

	db.Migrator().DropTable(&User{})
	db.Migrator().CreateTable(&User{}) //emailのunique keyがおかしくなるためテーブルを作り直す

	user := User{
		Email: sql.NullString{
			String: "a@a.com",
			Valid:  true,
		},
	}

	db.Create(&user)

	//db.AutoMigrate(&User{})
	/*
		user := User{
			Model: gorm.Model{
				CreatedAt: time.Now(),
			},
		}

		fmt.Println(user)
	*/
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
	/*
		user := User{
			Id:        1,
			FirstName: "Taiki updated",
			LastName:  "Noda updated",
			Email:     "hoge@hogehoge updated",
		}
		db.Updates(&user)
	*/

	//
	// delete user
	//
	/*
		user := User{
			Id: 1,
		}
		db.Delete(&user)
	*/

	//
	//ruery
	//
	/*
		users := []User{
			{
				FirstName: "Taiki",
				LastName:  "Noda",
				Email:     "taiki@hogehoge",
			},
			{
				FirstName: "Erika",
				LastName:  "Noma",
				Email:     "erika@hogehoge",
			},
			{
				FirstName: "Kazuaki",
				LastName:  "Ichitoku",
				Email:     "kazuaki@hogehoge",
			},
		}

		for _, user := range users {
			db.Create(&user)
		}
	*/
	/*
		user := User{}

		//db.First(&user)
		//db.Last(&user)
		db.Where("last_name", "Noda").First(&user)
		fmt.Println(user)
	*/
}
