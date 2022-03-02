package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/OybekNarzullaev/http-rest-api/internal/app/apiserver"
)

var (
	configPath string
)

// init config files
func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")	
}

func main() {
	// loading config files...
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatalf("Error occupired to connect config files: ", err)
	}

	// start server...
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}