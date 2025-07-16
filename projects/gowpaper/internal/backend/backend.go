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

func DetectBackend() *Backend {
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
