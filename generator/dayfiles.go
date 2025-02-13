package generator

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
)

// findLatestDay extracts the highest day number from the content
func (g *Generator) findLatestDay(content string) (int, error) {
	re := regexp.MustCompile(`"go-kata-machine/day(\d+)"`)
	matches := re.FindAllStringSubmatch(content, -1)

	latestDay := 0
	for _, match := range matches {
		if day, err := strconv.Atoi(match[1]); err == nil {
			if day > latestDay {
				latestDay = day
			}
		}
	}
	return latestDay, nil
}

// createDayFiles creates new directory and files for the day
func (g *Generator) createDayFiles(newDay int) error {
	newDayDir := fmt.Sprintf("day%d", newDay)
	if err := g.fs.MkdirAll(newDayDir, dirPerms); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	for filename, template := range templates {
		if err := g.createFile(newDayDir, filename, template, newDay); err != nil {
			return err
		}
	}
	return nil
}

// createFile creates a single file with the given template
func (g *Generator) createFile(dir, filename, template string, day int) error {
	path := filepath.Join(dir, filename)
	f, err := g.fs.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", path, err)
	}
	defer f.Close()

	content := fmt.Sprintf(template, day)
	if _, err := f.WriteString(content); err != nil {
		return fmt.Errorf("failed to write to file %s: %w", path, err)
	}
	return nil
}
