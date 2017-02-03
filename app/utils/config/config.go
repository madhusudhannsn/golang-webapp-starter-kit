package config

import (
	"os"

	"github.com/spf13/viper"
	"github.com/madhusudhannsn/go-web-app/app/utils/error"
	"github.com/madhusudhannsn/go-web-app/app/utils/logger"
)

//LoadConfig : This method loads the configuration
func LoadConfig() *error.AppError {
	logger.Info.Println("Initializing the configuration")
	fileName := os.Getenv("env")
	viper.SetConfigName(fileName)
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		return error.GetInstance("Config loading error", 1002, err)
	} else {
		return nil
	}

}
