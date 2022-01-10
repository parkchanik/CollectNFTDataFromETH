package config

import (
	"encoding/json"
	"os"
)

type ConnectConfig struct {
	URL string `json:"url"`
}

func LoadConfigration(path string) (ConnectConfig, error) {
	var config ConnectConfig
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&config)
	return config, nil
}
