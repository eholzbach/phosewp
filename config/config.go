// Package config returns the configuration data
package config

import (
	"flag"

	"github.com/BurntSushi/toml"
)

// Vars is a struct of all configuration options
type Vars struct {
	Channels      []string
	Coinmarketcap string
	Darksky       string
	Dbfile        string
	Dictionary    string
	Handle        string
	Newsapi       string
	Network       string
	Password      string
	SASL          bool
	TLS           bool
	Zipcodes      string
}

// Config reads the configuration file and returns a struct
func Config() (Vars, error) {
	var c Vars

	cpath := flag.String("config", "phosewp.toml", "configuration file")
	flag.Parse()

	_, err := toml.DecodeFile(*cpath, &c)
	return c, err
}
