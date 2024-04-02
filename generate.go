package main

import (
	"errors"
	"fmt"
	"go-kata-machine/scripts"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

const DirectoryPermission = 0750
const FilePermission = 0666

func getDayNumberValue(day string) (int, error) {
	dayNumberString := day[3:]
	dayNumberInt, err := strconv.Atoi(dayNumberString)
	if err != nil {
		return 0, errors.New("error getting day number")
	}
	return dayNumberInt, nil
}

func getMostRecentlyCreatedDay(rootDir string) int {
	files, err := os.ReadDir(rootDir)
	if err != nil {
		log.Fatal(err)
	}

	var dayFiles []os.DirEntry
	for _, file := range files {
		if file.IsDir() && strings.Contains(file.Name(), "day") {
			dayFiles = append(dayFiles, file)
		}
	}

	if len(dayFiles) == 0 {
		return 0
	}

	sortedDayFiles := append(dayFiles)
	sort.Slice(
		sortedDayFiles,
		func(i, j int) bool {
			ithNumber, _ := getDayNumberValue(dayFiles[i].Name())
			jthNumber, _ := getDayNumberValue(dayFiles[j].Name())
			return ithNumber > jthNumber
		},
	)

	recentDay, _ := getDayNumberValue(sortedDayFiles[0].Name())
	return recentDay
}

func main() {
	allDsaConfig := scripts.GetDsaDetails()
	rootDir, err := scripts.GetRootDir()

	if err != nil {
		log.Fatal(err)
	}

	recentlyCreatedDay := getMostRecentlyCreatedDay(rootDir)

	dayName := "day" + strconv.Itoa(recentlyCreatedDay+1)
	dayDirPath := path.Join(rootDir, dayName)

	// remove dir if it already exists
	err = os.RemoveAll(dayDirPath)
	err = os.Mkdir(dayDirPath, DirectoryPermission)
	if err != nil {
		log.Fatal(err)
	}

	for _, dsa := range scripts.GetDSA() {
		dsaDirPath := path.Join(dayDirPath, strings.ToLower(dsa))
		err = os.Mkdir(dsaDirPath, DirectoryPermission)

		dsaConfig, ok := allDsaConfig[dsa]
		if !ok {
			fmt.Printf("%s config not found", dsa)
			continue
		}

		if dsaConfig.DsaType == "fn" {
			err = createFunction(dsa, dsaDirPath, dsaConfig)
		}

		if err != nil {
			fmt.Println("error creating dsa file", err)
			continue
		}
	}

}

func createFunction(dsa, dsaDirPath string, dsaConfig scripts.DSADetails) error {
	text := `package %s
func %s(%s) %s {}
`

	formattedText := fmt.Sprintf(text, strings.ToLower(dsa), dsa, dsaConfig.ArgsTypes, dsaConfig.ReturnsType)
	err := os.WriteFile(dsaDirPath+"/"+dsa+".go", []byte(formattedText), FilePermission)
	if err != nil {
		return err
	}

	// create test for function
	dsaTest, ok := scripts.GetDsaTests()[dsa]
	if !ok {
		errInfo := fmt.Sprintf("%s test not found", dsa)
		return errors.New(errInfo)
	}

	err = os.WriteFile(dsaDirPath+"/"+dsa+"_test.go", []byte(dsaTest), FilePermission)

	return nil
}
