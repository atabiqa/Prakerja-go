package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		panic("Error init databse")
	}
	initMigrate()
}
