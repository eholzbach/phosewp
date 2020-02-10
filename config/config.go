// Package config returns the configuration data
package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type ConfigVars struct {
	Network  string
	Ssl      bool
	Handle   string
	Channels []string
	Darksky  string
	Dbfile   string
	Newsapi  string
}

// Config reads the configuration file and returns a struct of data
func Config() *ConfigVars {
	viper.SetConfigName("phosewp")
	viper.SetConfigName(".phosewp")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath("/usr/local/etc/")
	viper.AddConfigPath("$HOME/")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("config file not found, exiting")
		os.Exit(1)
	}

	viper.WatchConfig()

	a := &ConfigVars{
		Network:  viper.GetString("network"),
		Ssl:      viper.GetBool("ssl"),
		Handle:   viper.GetString("handle"),
		Channels: viper.GetStringSlice("channels"),
		Darksky:  viper.GetString("darksky"),
		Dbfile:   viper.GetString("dbfile"),
		Newsapi:  viper.GetString("newsapi"),
	}

	return a
}
