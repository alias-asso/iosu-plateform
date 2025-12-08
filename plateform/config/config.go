package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Sqlite struct {
	DbPath string
}

type Postgres struct {
	DbUrl    string
	Username string
	Password string
}

type Mysql struct {
	DbUrl    string
	Username string
	Password string
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

func DefaultConfig() Config {
	return Config{
		Sqlite: Sqlite{
			DbPath: fmt.Sprintf("/var/%s/db.sqlite", PlateformName),
		},
	}
}
