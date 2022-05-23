package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/models"
)

func readConfigs(logger *log.Logger) *models.Config {
	fileName := ""
	for i, a := range os.Args {
		if a == "-c" && len(os.Args) > i+1 {
			fileName = os.Args[i+1]
			break
		}
	}
	if fileName == "" {
		logger.Fatal("could not open config file")
		return nil
	}

	file, _ := os.Open(fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	var config models.Config
	err := decoder.Decode(&config)
	if err != nil {
		logger.Fatalf("was not able to load configs, err: %s", err.Error())
		return nil
	}
	return &config
}
