package database

import (
	"fmt"
	"log"

	"github.com/2ndsilencerz/cms-card-svc/configs"
	"github.com/2ndsilencerz/cms-card-svc/configs/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB the mysql
func InitDB() *gorm.DB {
	conf := configs.GetDatabaseConfig()
	db, err := gorm.Open(mysql.Open(conf), &gorm.Config{})
	if err != nil {
		msg := fmt.Sprintf("Error: %v\n", err)
		utils.LogToFile(msg)
		log.Panicln(msg)
	}
	return db
}

// CloseDB to close connection
func CloseDB(db *gorm.DB) {
	mysqlDB, err := db.DB()
	if err != nil {
		utils.LogToFile(fmt.Sprintf("Error: %v", err))
	}
	err = mysqlDB.Close()
	if err != nil {
		utils.LogToFile(fmt.Sprintf("Error: %v", err))
	}
}
