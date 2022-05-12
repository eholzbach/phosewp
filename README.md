phosewp
============

Yet another irc bot based on [go-ircevent](https://github.com/thoj/go-ircevent)

## Configuration

Configuration is supplied by the `--config` flag in [TOML](https://toml.io/en/).

Example:
```
network = 'chat.us.freenode.net:6697'
tls = true
handle = 'g1mpb0t'
channels = [ '#leetbotz', '#lamebotz' ]
db = '/path/to/sql.db'
accuweather = 'api_key'
zipcodes = '/path/to/zipcodes.json'
newsapi = 'api_key'
```