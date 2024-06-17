package main

import (
	"go-kata-machine/helpers"
	"log"
	"os"
	"strings"
)

func main() {
	// add options to use flags to declare with day directory to delete
	// selectDay := flag.String("day", "", "a string")
	// flag.Parse()
	// fmt.Println(len(*selectDay), *selectDay)

	rootDir, err := helpers.GetRootDir()
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
