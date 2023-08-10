package connection_db

import (
	"fmt"
	"log"
	"os"

	// model_bank "com.server/luis/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(){
	var err error

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local", "root", "root", "127.0.0.1", "3306", "bankgateway") 
	DB, err = gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect server", err.Error())
		os.Exit(1)
	}

	DB.Logger = logger.Default.LogMode(logger.Info)

	// this for migration
	// DB.Migrator(&model_bank.Bank{})

	fmt.Println("🚀 Connected Successfully to the Database")
}
