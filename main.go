package main

import (
	"flag"
	"fmt"
	"go-kata-machine/generator"
	"os"
)

type Config struct {
	clear    bool
	generate bool
	lastDay  bool
}

func main() {
	var config Config

	flag.BoolVar(&config.clear, "clear", false, "clear all day directories")
	flag.BoolVar(&config.generate, "generate", false, "generate new working day directory")
	flag.BoolVar(&config.lastDay, "lastDay", false, "get most recently created day directory")
	flag.Parse()

	if config.clear {
		clearDayDirectories()
		return
	}

	if config.generate {
		newGenerator := generator.NewGenerator(generator.RealFileSystem{})
		if err := newGenerator.Generate(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if config.lastDay {
		recentDay, err := getMostRecentlyCreatedDay()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(recentDay)
		}
		return
	}

}
