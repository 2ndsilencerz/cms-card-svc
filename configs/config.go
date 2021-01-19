package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func loadConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return v
}

// GetDatabaseConfig for database connection usage in database.go
func GetDatabaseConfig() string {
	v := loadConfig()
	return v.GetString("oracle")
}