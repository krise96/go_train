package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"gostartup/src/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "configPath", "src/configs/apiserver.toml", "path to a config file")
}

func main() {
	flag.Parse()

	config := apiserver.CreateConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	server := apiserver.ConfigureServer(config)

	if err := server.StartServer(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(server)
}
