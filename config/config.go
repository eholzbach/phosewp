// Package config returns the configuration data
package config

import (
	"github.com/spf13/viper"
)

type ConfigVars struct {
	Channels      []string
	Coinmarketcap string
	Darksky       string
	Dbfile        string
	Dictionary    string
	Handle        string
	Newsapi       string
	Network       string
	Password      string
	Sasl          bool
	Ssl           bool
	Zipcodes      string
}

// Config reads the configuration file and returns a struct of data
func Config() (*ConfigVars, error) {
	viper.SetConfigName("phosewp")
	viper.SetConfigName(".phosewp")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath("/usr/local/etc/")
	viper.AddConfigPath("$HOME/")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	viper.WatchConfig()

	a := &ConfigVars{
		Channels:      viper.GetStringSlice("channels"),
		Coinmarketcap: viper.GetString("coinmarketcap"),
		Darksky:       viper.GetString("darksky"),
		Dbfile:        viper.GetString("db"),
		Dictionary:    viper.GetString("dictionary"),
		Handle:        viper.GetString("handle"),
		Newsapi:       viper.GetString("newsapi"),
		Network:       viper.GetString("network"),
		Password:      viper.GetString("password"),
		Sasl:          viper.GetBool("sasl"),
		Ssl:           viper.GetBool("ssl"),
		Zipcodes:      viper.GetString("zipcodes"),
	}

	return a, err
}
