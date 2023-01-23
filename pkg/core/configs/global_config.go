package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetGlobalConfigString(profile string, key string, value string) interface{} {
	return viper.Get(fmt.Sprintf("%v.%v", profile, key))
}

func SetGlobalConfigString(profile string, key string, value interface{}) {
	viper.Set(fmt.Sprintf("%v.%v", profile, key), value)
	viper.WriteConfig()
}
