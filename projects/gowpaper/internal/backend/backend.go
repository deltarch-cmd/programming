package backend

import (
	"log"
	"os"
)

type Backend struct {
	Name string
	Cmd string
	Args []string
}

func detectBackend() *Backend {
	switch os.Getenv("XDG_SESSION_TYPE") {
	case "x11":
		return &Backend {
			Name: "feh",
			Cmd: "feh",
			Args: []string{"--bg-fill"},
		}
	case "wayland":
		return &Backend {
			Name: "swww",
			Cmd: "swww",
			Args: []string{"img"},
		}
	}
	log.Printf("Unknown session type")
	return nil
}

/* 
This function is mainly needed for full DE's like Plasma that use specific tools
We can argue that swww or feh are also usable, but I prefer the default tools for each env
*/
func SelectBackend() *Backend {
	switch os.Getenv("XDG_CURRENT_DESKTOP") {
	case "KDE":
		return &Backend{
			Name: "plasma-apply-wallpaperimage",
			Cmd: "plasma-apply-wallpaperimage",
			Args: []string{},
		}
	default:
		return detectBackend()
	}
}
