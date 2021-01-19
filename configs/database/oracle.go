package database

import (
	"fmt"
	"github.com/2ndSilencerz/cms-card-svc/config"
	"github.com/2ndSilencerz/cms-card-svc/config/utils"
	"github.com/cengsin/oracle"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {
	conf := config.GetDatabaseConfig()
	db, err := gorm.Open(oracle.Open(conf), &gorm.Config{})
	if err != nil {
		msg := fmt.Sprintf("Error: %v\n", err)
		utils.LogToFile(msg)
		log.Panicln(msg)
	}
	return db
}

func CloseDB(db *gorm.DB) {
	oracleDB, err := db.DB()
	if err != nil {
		utils.LogToFile(fmt.Sprintf("Error: %v", err))
	}
	err = oracleDB.Close()
	if err != nil {
		utils.LogToFile(fmt.Sprintf("Error: %v", err))
	}
}

