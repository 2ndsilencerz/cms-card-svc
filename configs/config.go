package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func loadConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error configs file: %s", err))
	}
	return v
}

// GetDatabaseConfig for database connection usage in database.go
func GetDatabaseConfig() string {
	v := loadConfig()
	// return v.GetString("oracle")
	return v.GetString("mysql")
}

// GetDefaultPort to define port for service
// can be defined either by passing port while executing or in var port in config.json
func GetDefaultPort() string {
	port := ":"
	if len(os.Args) > 1 && os.Args[1] != "" {
		port += os.Args[1]
	} else {
		v := loadConfig()
		port += v.GetString("port")
	}
	return port
}
