phosewp [![Build Status](https://travis-ci.org/eholzbach/phosewp.svg?branch=master)](https://travis-ci.org/eholzbach/phosewp)
============

Yet another irc bot based on [go-ircevent](https://github.com/thoj/go-ircevent)

## Configuration

This uses [viper](https://github.com/spf13/viper) to resolve configuration files. JSON, TOML, YAML, and HCL are valid formats.

Example:
```
network: chat.us.freenode.net:6697
ssl: true
handle: g1mpb0t
channels:
  - #leetbotz
darksky: api_key
zipcodes: /path/to/zipcodes.json
dbfile: /path/to/sql.db
newsapi: api_key
```
