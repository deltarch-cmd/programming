package main

import (
	"gowpaper/internal/backend"
	"gowpaper/internal/config"
	"gowpaper/internal/imagepicker"

	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

func setWallpaper(b *backend.Backend, path, picture string) {
	wallpaper := filepath.Join(path, picture)
	cmd := exec.Command(b.Cmd, append(b.Args, wallpaper)...)

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to set wallpaper: %v", err)
	}
	fmt.Println("Wallpaper set successfully!")
}

func main() {
	conf := config.Load()

	// Check the backend and choose a tool to set the wallpaper
	tooling := backend.DetectBackend()
	if tooling == nil {
		log.Fatal("No backend detected")
	}

	// Check pictures from the selected directory
	// !TODO Make it so the directory can be choosen if none is saved
	var fullPath string = filepath.Join(conf.Home, conf.LastDirUsed)
	var pictures []string = imagepicker.GetPicturesFromDir(fullPath, conf.AllowedExt)

	if len(pictures) == 0 {
		log.Fatal("No valid pictures found.")
	}

	var imageSelected string = imagepicker.ChooseWallpaper(pictures)
	setWallpaper(tooling, fullPath, imageSelected)

	conf.LastWallpaper = imageSelected
	if err := config.Save("config.json", &conf); err != nil {
		log.Printf("Failed to update config %v", err)
	}
}
