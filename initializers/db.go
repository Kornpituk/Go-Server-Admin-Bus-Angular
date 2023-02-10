package initializers

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToBD() {
	var err error

	dsn := "host=localhost user=postgres password=1234 dbname=admin-bus port=5050 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to open database")
	} else {
		fmt.Printf("Database created successfully")
	}
}
