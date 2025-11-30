package main

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Sqlite struct {
	DbPath string
}

type Config struct {
	Sqlite Sqlite
}

func ParseConfig(path string) (error, Config) {
	var config Config
	tomlData, err := os.ReadFile(path)
	if err != nil {
		return err, Config{}
	}
	_, err = toml.Decode(string(tomlData), config)
	if err != nil {
		return err, Config{}
	}
	return nil, config
}
