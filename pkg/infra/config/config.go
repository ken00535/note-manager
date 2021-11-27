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

type config struct{}

// Init config
func Init() Config {
	path, exist := os.LookupEnv("CONFIG_PATH")
	if !exist {
		path = "./config"
	}
	once.Do(func() {
		log = logger.New()
	})
	setting = viper.New()
	setting.SetConfigName("config") // name of config file (without extension)
	setting.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	setting.AddConfigPath(path)     // optionally look for config in the working directory
	err := setting.ReadInConfig()   // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		log.Panic(err)
	}
	return &config{}
}

// GetDbAddress from config
func (*config) GetDbAddress() string {
	return setting.GetString("server_address")
}

// GetDbPort from config
func (*config) GetDbPort() int {
	return setting.GetInt("server_port")
}

// GetRdbAdress from config
func (*config) GetRdbAdress() string {
	return setting.GetString("rdb_address")
}

// GetRdbPort from config
func (*config) GetRdbPort() int {
	return setting.GetInt("rdb_port")
}

// GetRdbPassword from config
func (*config) GetRdbPassword() string {
	return setting.GetString("rdb_password")
}

// GetUsername from config
func (*config) GetUsername() string {
	return setting.GetString("service_account")
}

// GetPassword from config
func (*config) GetPassword() string {
	return setting.GetString("service_password")
}

// GetSecret from config
func (*config) GetSecret() string {
	return setting.GetString("service_jwt_secert")
}
