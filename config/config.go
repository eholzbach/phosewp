// Package config returns the configuration data
package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type ConfigVars struct {
	Network    string
	Sasl       bool
	Ssl        bool
	Handle     string
	Password   string
	Channels   []string
	Darksky    string
	Zipcodes   string
	Dbfile     string
	Newsapi    string
	Dictionary string
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
		Network:    viper.GetString("network"),
		Sasl:       viper.GetBool("sasl"),
		Ssl:        viper.GetBool("ssl"),
		Handle:     viper.GetString("handle"),
		Password:   viper.GetString("password"),
		Channels:   viper.GetStringSlice("channels"),
		Darksky:    viper.GetString("darksky"),
		Zipcodes:   viper.GetString("zipcodes"),
		Dbfile:     viper.GetString("dbfile"),
		Newsapi:    viper.GetString("newsapi"),
		Dictionary: viper.GetString("dictionary"),
	}

	return a
}
