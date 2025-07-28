package config

import (
	"github.com/kirsle/configdir"
	"github.com/spf13/viper"
)

func ResolveAlias(configKey string, alias string) string {
	aliasMap := viper.GetStringMapString(configKey)
	if value, ok := aliasMap[alias]; ok {
		return value
	} else {
		return alias
	}
}

func InitConfig() error {
	confDir := configdir.LocalConfig("ptt")
	viper.SetDefault("color", false)
	viper.SetDefault("header", false)
	viper.SetDefault("timestamp", false)
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(confDir)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
