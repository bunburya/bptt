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

func ResolveAliases(configKey string, aliases []string) []string {
	aliasMap := viper.GetStringMapString(configKey)
	var resolved []string
	for _, alias := range aliases {
		if value, ok := aliasMap[alias]; ok {
			resolved = append(resolved, value)
		} else {
			resolved = append(resolved, alias)
		}
	}
	return resolved
}

func InitConfig() {
	viper.SetDefault("color", false)
	viper.SetDefault("header", false)
	viper.SetDefault("timestamp", false)
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	confFile := viper.GetString("config")
	if confFile != "" {
		viper.SetConfigFile(confFile)
	} else {
		confDir := configdir.LocalConfig("ptt")
		viper.AddConfigPath(confDir)
	}
	_ = viper.ReadInConfig()
}
