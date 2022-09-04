package main

import (
	"fmt"
	"log"

	"unpacking/internal/app"
	"unpacking/internal/config"
	"unpacking/internal/reader"
)

func main() {
	confCore, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Config err: %v", err)
	}
	conf, err := config.ParseConfig(confCore)
	if err != nil {
		log.Fatalf("Config err: %v", err)
	}

	app := app.NewApp(reader.FileReader{}, conf)
	result, err := app.Execute()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(result)
}
