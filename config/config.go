package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func Config() (string, bool, string, []string, string, string) {

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

	c := viper.GetStringSlice("channels")
	n := c[:0]
	for _, v := range c {
		if strings.HasPrefix(v, "#") == false {
			v = "#" + v
		}

		n = append(n, v)

	}

	return viper.GetString("network"), viper.GetBool("ssl"), viper.GetString("handle"), n, viper.GetString("darksky"), viper.GetString("newsapi")
}
