package imagepicker

import (
	"os"
	"fmt"
	"log"
	"strings"
)

func GetPicturesFromDir(path string, allowedExts []string) []string {
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

func ChooseWallpaper(pictures []string) string {
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
