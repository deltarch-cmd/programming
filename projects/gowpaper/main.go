package main

import (
	"gowpaper/internal/backend"
	"gowpaper/internal/config"
	"gowpaper/internal/imagepicker"
	"os"

	"fmt"
	"log"
	"os/exec"
	"path/filepath"

	"flag"
	"math/rand/v2"
)

func setWallpaper(b *backend.Backend, path, picture string) {
	wallpaper := filepath.Join(path, picture)
	cmd := exec.Command(b.Cmd, append(b.Args, wallpaper)...)

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to set wallpaper: %v", err)
	}
	fmt.Println("Wallpaper " + picture + " set successfully!")
}

func main() {
	conf := config.Load()

	// Flags
	restoreFlag := flag.Bool("restore", false, "Restore last used wallpaper")
	randomFlag := flag.Bool("random", false, "Choose a random wallpaper")
	flag.Parse()

	// Check the backend and choose a tool to set the wallpaper
	tooling := backend.SelectBackend()
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

	if *restoreFlag {
		log.Printf("Restore flag activated, restoring last used wallpaper...")
		setWallpaper(tooling, fullPath, conf.LastWallpaper)
		os.Exit(0)
	}

	if *randomFlag {
		log.Printf("Random flag activated, setting a random wallpaper...")
		randomN := rand.IntN(len(pictures))
		setWallpaper(tooling, fullPath, pictures[randomN + 1])
		os.Exit(0)
	}

	var imageSelected string = imagepicker.ChooseWallpaper(pictures)
	setWallpaper(tooling, fullPath, imageSelected)

	conf.LastWallpaper = imageSelected
	if err := config.Save("config.json", &conf); err != nil {
		log.Printf("Failed to update config %v", err)
	}
}
