package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func generate() error {
	// 1. Read the current registry file
	registryPath := "tests/registry.go"
	content, err := os.ReadFile(registryPath)
	if err != nil {
		return fmt.Errorf("failed to read registry file: %v", err)
	}

	// 2. Find the latest day number
	re := regexp.MustCompile(`"go-kata-machine/day(\d+)"`)
	matches := re.FindAllStringSubmatch(string(content), -1)

	latestDay := 0
	for _, match := range matches {
		if day, err := strconv.Atoi(match[1]); err == nil {
			if day > latestDay {
				latestDay = day
			}
		}
	}

	newDay := latestDay + 1

	// 3. Create new import statement
	newImport := fmt.Sprintf("\t\"go-kata-machine/day%d\"", newDay)

	// 4. Add new import to imports block
	importBlock := regexp.MustCompile(`import \(([\s\S]*?)\)`)
	newContent := importBlock.ReplaceAllStringFunc(string(content), func(match string) string {
		return strings.TrimSuffix(match, ")") + newImport + "\n)"
	})

	// 5. Add new registry entry
	registryEntry := fmt.Sprintf(`	"day%d": {
		binarySearch: day%d.BinarySearchList,
		bubbleSort:   day%d.BubbleSort,
	},`, newDay, newDay, newDay)

	// Find Registry map and add new entry
	registryBlock := regexp.MustCompile(`var Registry = map\[string\]Day\{([\s\S]*?)\}`)
	newContent = registryBlock.ReplaceAllStringFunc(newContent, func(match string) string {
		return strings.TrimSuffix(match, "}") + registryEntry + "\n}"
	})

	// 6. Create the new day directory and files
	newDayDir := fmt.Sprintf("day%d", newDay)
	err = os.MkdirAll(newDayDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// Create template files
	templates := map[string]string{
		"binary_search.go": `package day%d

func BinarySearchList(haystack []int, needle int) bool {
	// TODO: implement
	return false
}
`,
		"bubble_sort.go": `package day%d

func BubbleSort(arr []int) []int {
	// TODO: implement
	return arr
}
`,
	}

	for filename, template := range templates {
		filepath := filepath.Join(newDayDir, filename)
		f, err := os.Create(filepath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %v", filepath, err)
		}
		defer f.Close()

		content := fmt.Sprintf(template, newDay)
		_, err = f.WriteString(content)
		if err != nil {
			return fmt.Errorf("failed to write to file %s: %v", filepath, err)
		}
	}

	// 7. Write the modified content back to registry.go
	err = os.WriteFile(registryPath, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write registry file: %v", err)
	}

	fmt.Printf("Successfully generated day%d\n", newDay)
	return nil
}

func main() {
	if err := generate(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
