package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetModuleName reads the go.mod file in the current directory and extracts the module name.
func GetModuleName(projectPath string) (string, error) {
	file, err := os.Open(filepath.Join(projectPath, "go.mod"))
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			// Extract the module name from the line.
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				return parts[1], nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("module directive not found in go.mod")
}
