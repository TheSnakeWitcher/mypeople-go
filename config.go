package main

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"

	"errors"
	"os"
)

const (
	ConfigPathDev    string = "/home/mr-papi/SoftwareCode/Projects/mypeople"
	ConfigPathSystem string = "/etc/mypeople/"
	ConfigPathUser   string = "/home/mr-papi/.config/mypeople"
	ConfigFileName   string = "mypeople-config"
	ConfigFileType   string = "toml"
)

type configuration struct {
	LogFile      string `mapstructure:LogFile`
	DbDriver     string `mapstructure:DbDriver`
	DbSource     string `mapstructure:DbSource`
	DbSourceMem  string `mapstructure:DbSourceMem`
	DbSourceDev string `mapstructure:DbSourceDev`
}

var (
	Config  configuration
	LogFile *os.File
	Logger  zerolog.Logger
)

func InitConfig() (err error) {
	viper.SetConfigName(ConfigFileName)
	viper.SetConfigType(ConfigFileType)
	viper.AddConfigPath(ConfigPathDev)
	viper.AddConfigPath(ConfigPathSystem)
	viper.AddConfigPath(ConfigPathUser)

	err = viper.ReadInConfig()
	if err != nil {
		panic(errors.New("failed to load config file"))
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(errors.New("failed to save config"))
	}

	return nil
}

func InitLogger() (err error) {
	LogFile, err = os.OpenFile(Config.LogFile, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(errors.New("failed to open logfile"))
	}

	fileWriter := zerolog.New(LogFile).With().Logger()
	Logger = zerolog.New(fileWriter).With().Timestamp().Logger()
	return nil
}
