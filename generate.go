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
	rootDir, err := scripts.GetRootDir()

	if err != nil {
		log.Fatal(err)
	}

	recentlyCreatedDay := getMostRecentlyCreatedDay(rootDir)

	dayName := "day" + strconv.Itoa(recentlyCreatedDay+1)
	dayDirPath := path.Join(rootDir, dayName)

	// remove dir if it already exists then create day directory
	err = os.RemoveAll(dayDirPath)
	err = os.Mkdir(dayDirPath, DirectoryPermission)
	if err != nil {
		log.Fatal(err)
	}

	for _, dsa := range scripts.GetDSA() {
		// create algo kata directory
		dsaDirPath := path.Join(dayDirPath, strings.ToLower(dsa))
		err = os.Mkdir(dsaDirPath, DirectoryPermission)

		dsaConfig, ok := scripts.GetDsaDetails()[dsa]
		if !ok {
			fmt.Printf("%s config not found", dsa)
			continue
		}

		switch dsaConfig.DsaType {
		case "struct":
			err = createStruct(dsa, dsaDirPath, dsaConfig)
			break
		default:
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

	// create test file and content
	err = createDsaTest(dsa, dsaDirPath, err)
	if err != nil {
		return err
	}

	return nil
}

func createStruct(dsa, dsaDirPath string, dsaConfig scripts.DSADetails) error {
	structDefinition := `type Node[T any] struct {
	value T
	// pointer-value of nextNode, zerothValue is nil
	prev *Node[T]
}

type %s[T any] struct {
	Length int
	// reference to pointer-value of headNode, zerothValue is nil
	head *Node[T]
}`

	formattedStructDefinition := fmt.Sprintf(structDefinition, dsa)

	methodsDefinition := ""
	methodsDeclaration := ""

	for _, method := range dsaConfig.Methods {
		definitionText := fmt.Sprintf("%s(%s) %s\n", method.Name, method.ArgsTypes, method.ReturnTypes)
		methodsDefinition += definitionText

		declarationText := fmt.Sprintf(
			"func(%s *%s[%s]) %s(%s) %s {}\n",
			strings.ToLower(dsa[:1]), dsa, dsaConfig.Generic, method.Name, method.ArgsTypes, method.ReturnTypes,
		)
		methodsDeclaration += declarationText
	}

	text := `package %s

%s

type Methods[%s any] interface {
%s
}

%s
`

	formattedText := fmt.Sprintf(
		text,
		strings.ToLower(dsa), formattedStructDefinition, dsaConfig.Generic, methodsDefinition, methodsDeclaration,
	)

	err := os.WriteFile(dsaDirPath+"/"+dsa+".go", []byte(formattedText), FilePermission)
	if err != nil {
		return err
	}

	// create test file and content
	err = createDsaTest(dsa, dsaDirPath, err)
	if err != nil {
		return err
	}

	return nil

}

func createDsaTest(dsa string, dsaDirPath string, err error) error {
	dsaTest, ok := scripts.GetDsaTests()[dsa]
	if !ok {
		errInfo := fmt.Sprintf("%s test not found", dsa)
		return errors.New(errInfo)
	}
	err = os.WriteFile(dsaDirPath+"/"+dsa+"_test.go", []byte(dsaTest), FilePermission)
	if err != nil {
		return err
	}
	return nil
}
