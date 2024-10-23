package scripts

import (
	"errors"
	"fmt"
	"go-kata-machine/helpers"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

const (
	DirectoryPermission = 0750
	FilePermission      = 0666
)

func getDayNumberValue(day string) (int, error) {
	dayNumberString := day[3:]
	dayNumberInt, err := strconv.Atoi(dayNumberString)
	if err != nil {
		return 0, errors.New("error getting day number")
	}
	return dayNumberInt, nil
}

func GetMostRecentlyCreatedDay() int {
	rootDir, err := helpers.GetRootDir()
	if err != nil {
		log.Fatal(err)
	}

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

func GenerateNewDayDirectory() {
	rootDir, err := helpers.GetRootDir()

	if err != nil {
		log.Fatal(err)
	}

	recentlyCreatedDay := GetMostRecentlyCreatedDay()

	dayName := "day" + strconv.Itoa(recentlyCreatedDay+1)
	dayDirPath := path.Join(rootDir, dayName)

	// remove dir if it already exists then create day directory
	err = os.RemoveAll(dayDirPath)
	err = os.Mkdir(dayDirPath, DirectoryPermission)
	if err != nil {
		log.Fatal(err)
	}

	for dsa, dsaConfig := range helpers.GetDsaDetails() {
		// create algo kata directory e.g .day1/Stack/, .day1/Queue/
		dsaDirPath := path.Join(dayDirPath, strings.ToLower(dsa))
		err = os.Mkdir(dsaDirPath, DirectoryPermission)

		// DsaType: either fn or struct atm
		switch dsaConfig.DsaType {
		case "struct":
			err = createStruct(dsa, dsaDirPath, dsaConfig)
			break
		default:
			err = createFunction(dsa, dsaDirPath, dsaConfig)
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		err = createDsaTest(dsa, dsaDirPath, rootDir)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func createFunction(dsa, dsaDirPath string, dsaConfig helpers.DSADetails) error {
	text := `package %s
func %s(%s) %s {}
`

	formattedText := fmt.Sprintf(text, strings.ToLower(dsa), dsa, dsaConfig.ArgsTypes, dsaConfig.ReturnsType)
	err := os.WriteFile(dsaDirPath+"/"+dsa+".go", []byte(formattedText), FilePermission)
	if err != nil {
		return err
	}
	return nil
}

func createStruct(dsa, dsaDirPath string, dsaConfig helpers.DSADetails) error {
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
	return nil
}

func createDsaTest(dsa, dsaDirPath, rootDir string) error {
	testTxtPath := rootDir + "/test-txts/" + dsa + ".txt"
	dsaTest, err := os.ReadFile(testTxtPath)
	if err != nil {
		errInfo := fmt.Sprintf("%s test not found", dsa)
		return errors.New(errInfo)
		// panic(err)
	}
	err = os.WriteFile(dsaDirPath+"/"+dsa+"_test.go", dsaTest, FilePermission)
	if err != nil {
		return err
	}
	return nil
}
