package config

import (
	"log"
	"os"
	"encoding/json"
)

type Config struct {
	Home string `json:"HOME"`
	LastDirUsed string `json:"last_directory_used"`
	LastWallpaper string `json:"last_wallpaper"`
	AllowedExt []string `json:"allowed_extensions"`
}

// Load the JSON file and return it
func Load() Config {
	content, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error during Unmarshal ono: ", err)
	}
	return config
}

// Save the JSON file with the given data
func Save(configPath string, conf *Config) error {
	data, err := json.MarshalIndent(conf, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0644)
}
