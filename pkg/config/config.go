package config

import (
	"note-manager/pkg/infra/logger"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	setting *viper.Viper
	db      *viper.Viper
	log     logger.Logger
	once    sync.Once
)

// Init config
func Init(logInst logger.Logger) {
	path, exist := os.LookupEnv("CONFIG_PATH")
	if !exist {
		path = "./config"
	}
	log = logInst
	setting = viper.New()
	setting.SetConfigName("config") // name of config file (without extension)
	setting.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	setting.AddConfigPath(path)     // optionally look for config in the working directory
	err := setting.ReadInConfig()   // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		log.Panic(err)
	}
}

// GetDbAddress from config
func GetDbAddress() string {
	return setting.GetString("server_address")
}

// GetDbPort from config
func GetDbPort() int {
	return setting.GetInt("server_port")
}

// GetRdbAdress from config
func GetRdbAdress() string {
	return setting.GetString("rdb_address")
}

// GetRdbPort from config
func GetRdbPort() int {
	return setting.GetInt("rdb_port")
}

// GetRdbPassword from config
func GetRdbPassword() string {
	return setting.GetString("rdb_password")
}
