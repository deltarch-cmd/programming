package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Config struct {
	Home string `json:"HOME"`
	LastDirUsed string `json:"last_directory_used"`
	LastWallpaper string `json:"last_wallpaper"`
	AllowedExt []string `json:"allowed_extensions"`
}

func getData() Config {
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

func getPicturesFromDir(path string, allowedExts []string) []string {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var pictures []string
	for _, entry := range entries {
		name := entry.Name()
		for _, ext := range allowedExts {
			if strings.HasSuffix(strings.ToLower(name), strings.ToLower(ext)) {
				pictures = append(pictures, name)
				break
			}
		}
	}

	return pictures
}

func chooseWallpaper(pictures []string) string {
	for i, pic := range pictures {
		fmt.Printf("%d) %s\n", i+1, pic)
	}

	var election int
	fmt.Scanf("%d", &election)

	if election < 1 || election > len(pictures) {
		log.Fatalf("Invalid selection: %d", election)
	}

	return pictures[election-1]
}

func main() {
	var conf Config = getData()

	var full_path string = filepath.Join(conf.Home, conf.LastDirUsed)
	var pictures []string = getPicturesFromDir(filepath.Join(full_path), conf.AllowedExt)

	if len(pictures) == 0 {
		log.Fatal("No valid pictures found.")
	}

	var selected string = chooseWallpaper(pictures)
	fmt.Printf("The wallpaper selected is %s, lets try setting it!\n", selected)

	err := exec.Command("swww", "img", filepath.Join(full_path, selected)).Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Wallpaper should be set! o7")
}
