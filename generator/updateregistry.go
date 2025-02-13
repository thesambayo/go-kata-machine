package generator

import (
	"fmt"
	"regexp"
	"strings"
)

// updateRegistryImports adds new import statement to the imports block
func (*Generator) updateRegistryImports(content string, newDay int) string {
	newImport := fmt.Sprintf("\t\"go-kata-machine/day%d\"", newDay)
	importBlock := regexp.MustCompile(`import \(([\s\S]*?)\)`)

	return importBlock.ReplaceAllStringFunc(content, func(match string) string {
		return strings.TrimSuffix(match, ")") + newImport + "\n)"
	})
}

// updateRegistry adds new day to the content of the registry map
func (*Generator) updateRegistry(content string, newDay int) string {
	// Create the new registry entry with proper formatting
	// %d is replaced with newDay value
	registryEntry := fmt.Sprintf(`	"day%d": {
		binarySearch: day%d.BinarySearchList,
		bubbleSort:   day%d.BubbleSort,
	}`, newDay, newDay, newDay)

	// Regular expression breakdown:
	// (var Registry = map\[string\]Day\{) 	- Group 1: Captures the map declaration
	// ([\s\S]*?)                         	- Group 2: Captures all existing entries (lazy match)
	// (\})                               	- Group 3: Captures the closing brace
	registryBlock := regexp.MustCompile(`(var Registry = map\[string\]Day\{)([\s\S]*?)(\})`)

	// ReplaceAllStringFunc replaces all matches using a custom function
	return registryBlock.ReplaceAllStringFunc(content, func(match string) string {
		// FindStringSubmatch returns all groups captured in the regex
		// submatches[0] = full match
		// submatches[1] = map declaration
		// submatches[2] = existing entries
		// submatches[3] = closing brace
		submatches := registryBlock.FindStringSubmatch(match)

		// Safety check: ensure we have all required groups
		if len(submatches) != 4 {
			// If regex pattern doesn't match as expected, return original content
			fmt.Printf("Warning: Registry format not recognized\n")
			return match
		}

		// Extract parts for better readability
		declaration := submatches[1]     // "var Registry = map[string]Day{"
		existingEntries := submatches[2] // All current map entries

		// Handle comma separation
		// If there are existing entries (ignoring whitespace), add a comma
		trimmedEntries := strings.TrimSpace(existingEntries)
		if trimmedEntries != "" {
			registryEntry += ","
		}

		result := declaration + "\n" + registryEntry + existingEntries + "\n}"

		return result
	})
}

// validateRegistryFormat performs basic validation on the registry format
func (*Generator) validateRegistryFormat(content string) bool {
	// Check for basic structural elements
	checks := []struct {
		element string
		error   string
	}{
		{`map[string]Day`, "Missing map declaration"},
		{`binarySearch:`, "Missing binarySearch field"},
		{`bubbleSort:`, "Missing bubbleSort field"},
	}

	for _, check := range checks {
		if !strings.Contains(content, check.element) {
			fmt.Printf("Validation Error: %s\n", check.error)
			return false
		}
	}

	// Check for balanced braces
	braceCount := 0
	for _, char := range content {
		switch char {
		case '{':
			braceCount++
		case '}':
			braceCount--
		}
		if braceCount < 0 {
			fmt.Printf("Validation Error: Unbalanced braces\n")
			return false
		}
	}
	if braceCount != 0 {
		fmt.Printf("Validation Error: Unbalanced braces\n")
		return false
	}

	return true
}

// Optional: Add helper method to format the registry entry
func (*Generator) formatRegistryEntry(day int) string {
	return fmt.Sprintf(`	"day%d": {
		binarySearch: day%d.BinarySearchList,
		bubbleSort:   day%d.BubbleSort,
	}`, day, day, day)
}
