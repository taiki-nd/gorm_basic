package main

import (
	"fmt"
	"gorm_basic/config"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
)

func main() {
	fmt.Println(config.Config.DbUserDevelop)
	fmt.Println(config.Config.DbPasswordDevelop)
}
