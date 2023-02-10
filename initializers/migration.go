package initializers

import (
	"server/model"
)

func Migration(){
	if DB != nil {
		DB.AutoMigrate(&model.User{})
	}
}