package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/GuillermoMarcel/financial-chat/internal/chat-bot/config"
)

func readConfigs(logger *log.Logger) *config.Configs {
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
	var config config.Configs
	err := decoder.Decode(&config)
	if err != nil {
		logger.Fatalf("was not able to load configs, err: %s", err.Error())
		return nil
	}
	return &config
}
