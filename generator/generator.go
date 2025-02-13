package generator

import (
	"fmt"
	"os"
)

// Configuration constants
const (
	registryPath = "tests/registry.go"
	filePerms    = 0644
	dirPerms     = 0755
)

// FileSystem interface for better testing and dependency injection
type FileSystem interface {
	ReadFile(string) ([]byte, error)
	WriteFile(string, []byte, os.FileMode) error
	MkdirAll(string, os.FileMode) error
	Create(string) (*os.File, error)
}

// RealFileSystem implements FileSystem interface using actual OS calls
type RealFileSystem struct{}

func (fs RealFileSystem) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (fs RealFileSystem) WriteFile(path string, data []byte, perm os.FileMode) error {
	return os.WriteFile(path, data, perm)
}

func (fs RealFileSystem) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (fs RealFileSystem) Create(path string) (*os.File, error) {
	return os.Create(path)
}

// Generator handles the generation of new day files
type Generator struct {
	fs FileSystem
}

// NewGenerator creates a new Generator instance
func NewGenerator(fs FileSystem) *Generator {
	return &Generator{fs: fs}
}

// Generate handles the entire generation process
func (g *Generator) Generate() error {
	// Read registry file
	content, err := g.fs.ReadFile(registryPath)
	if err != nil {
		return fmt.Errorf("failed to read registry file: %w", err)
	}

	// Find latest day and increment
	latestDay, err := g.findLatestDay(string(content))
	if err != nil {
		return fmt.Errorf("failed to find latest day: %w", err)
	}
	newDay := latestDay + 1

	// Update content with new imports
	newContent := g.updateRegistryImports(string(content), newDay)
	// Update content with registry entries
	newContent = g.updateRegistry(newContent, newDay)

	// Validate the final content before creating new files and updating registry file
	if !g.validateRegistryFormat(newContent) {
		return fmt.Errorf("generated content validation failed")
	}

	// Create new files
	if err := g.createDayFiles(newDay); err != nil {
		return err
	}

	// Write updated registry file
	if err := g.fs.WriteFile(registryPath, []byte(newContent), filePerms); err != nil {
		return fmt.Errorf("failed to write registry file: %w", err)
	}

	fmt.Printf("Successfully generated day%d\n", newDay)
	return nil
}
