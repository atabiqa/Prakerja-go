package configs

import "prakerja11/controllers/user"

func initMigrate() {
	DB.AutoMigrate(&{})
}
