package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const PlateformName = "iosu"

var configDirPath string = fmt.Sprintf("/etc/%s", PlateformName)
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

}
