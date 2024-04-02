package main

import (
	"go-kata-machine/scripts"
	"log"
	"os"
	"strings"
)

func main() {
	rootDir, err := scripts.GetRootDir()
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir(rootDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() && strings.Contains(file.Name(), "day") {
			err = os.RemoveAll(file.Name())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
