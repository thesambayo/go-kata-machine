package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func getRootDir() (string, error) {
	_, runtimeFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to find root directory")
	}
	fmt.Println(filepath.Dir(runtimeFile))
	return filepath.Dir(runtimeFile), nil
}

func clearDayDirectories() {
	rootDir, err := getRootDir()
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

func getMostRecentlyCreatedDay() (string, error) {
	rootDir, err := getRootDir()
	if err != nil {
		return "", err
	}

	files, err := os.ReadDir(rootDir)
	if err != nil {
		return "", err
	}

	maxDay := 0
	dayPrefix := "day"

	for _, file := range files {
		if file.IsDir() && strings.HasPrefix(file.Name(), dayPrefix) {
			dayNum := strings.TrimPrefix(file.Name(), dayPrefix)
			dayInt := 0
			_, err := fmt.Sscanf(dayNum, "%d", &dayInt)
			if err == nil && dayInt > maxDay {
				maxDay = dayInt
			}
		}
	}

	if maxDay == 0 {
		return "", errors.New("no day directories found")
	}

	return fmt.Sprintf("day%d", maxDay), nil
}
