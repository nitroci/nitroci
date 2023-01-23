package configs

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func EnsureConfiguration(configFile string) (string, error) {
	configHome := filepath.Dir(configFile)
	configName := filepath.Base(configFile)
	configType := filepath.Ext(configFile)
	configPath := filepath.Join(configHome, configName)
	viper.AddConfigPath(configHome)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	abs, _ := filepath.Abs(configHome)
	if _, err := os.Stat(abs); os.IsNotExist(err) {
		os.MkdirAll(configHome, 0700)
	}
	_, err := os.Stat(configPath)
	if err!= nil && !os.IsExist(err) { 
		if err := viper.SafeWriteConfig(); err != nil {
			return "", err
		}
	}
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return "", err
	}
	return configFile, nil
}
