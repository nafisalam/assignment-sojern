package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	LogsPath       string `yaml:"logsPath"`
	ImagePath      string `yaml:"imagePath"`
	BannerPath     string `yaml:"bannerPath"`
	StatusFilePath string `yaml:"statusFilePath"`
}

// InitConfig to initialize the config
func InitConfig(path string) Config {
	os.Setenv("env", "e0")
	env := os.Getenv("env")

	viper.SetConfigName(env)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.GetViper().SetEnvKeyReplacer(strings.NewReplacer(".", "-"))
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error to read the config file ", err)

	}
	var config Config
	err = viper.Unmarshal(&config)

	if err != nil {
		log.Fatal("Error to unmarshal the config file", err)

	}
	return config
}
