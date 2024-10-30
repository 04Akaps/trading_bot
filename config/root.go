package config

import (
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	Slack          Slack
	CryptoCurrency map[string]CryptoCurrency
}

func NewCfg(file string) Config {
	c := new(Config)

	if f, err := os.Open(file); err != nil {
		panic(err)
	} else {
		if err = toml.NewDecoder(f).Decode(c); err != nil {
			panic(err)
		} else {
			return *c
		}
	}
}
