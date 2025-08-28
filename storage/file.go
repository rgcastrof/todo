package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

func createDir() (string, error) {
	confPath, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("Failed to read config dir: %w", err)
	}
	todoPath := filepath.Join(confPath, "todo")
	err = os.MkdirAll(todoPath, 0750)
	if err != nil {
		return "", fmt.Errorf("Failed to create todo config dir: %w", err)
	}
	return todoPath, nil
}

func SaveTasks(data []byte) error {
	file, err := os.Create("tasks.json")
	if err != nil {
		return fmt.Errorf("Failed to create json file: %w", err)
	}
	defer file.Close()

	return os.WriteFile("tasks.json", data, 0644)
}

func LoadTasks() (string, error) {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return "", fmt.Errorf("Failed to read file: %w", err)
	}
	return string(data), nil
}
