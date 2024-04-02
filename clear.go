package main

import (
	"fmt"
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
		info, err := file.Info()

		if err != nil {
			log.Fatal(err)
		}

		if file.IsDir() && strings.Contains(file.Name(), "day") {
			fmt.Println(file.Name(), info.Size())
			err = os.Remove(file.Name())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
