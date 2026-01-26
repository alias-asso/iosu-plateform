package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/alias-asso/iosu/internal/config"
	"github.com/alias-asso/iosu/internal/server"
)

var configDirPath string = fmt.Sprintf("/etc/%s", config.PlateformName)
var configPath = flag.String("c", filepath.Join(configDirPath, "config.toml"), "Config file path.")

func main() {
	flag.Parse()

	// Check if config file exist and is readable
	if _, err := os.Stat(*configPath); err != nil {
		if os.IsNotExist(err) {
			log.Fatalln("Error : config file not found.")
		} else {
			log.Fatalln("Error : unable to read config file.")
		}
	}

	config, err := config.ParseConfig(*configPath)
	if err != nil {
		log.Fatalln("Error parsing config : " + err.Error())
	}

	serv, err := server.NewServer(config)
	if err != nil {
		log.Fatalln("Error creating server : " + err.Error())
	}

	err = serv.SetupServer(config)
	if err != nil {
		log.Fatalln("Error setting up server : " + err.Error())
	}

	serv.Start(config.ServerPort)
}
