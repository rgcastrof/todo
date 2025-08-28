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

func getFilePath() (string, error) {
	dir, err := createDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "tasks.json"), nil
}

func SaveTasks(data []byte) error {
	filepath, err := getFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		return fmt.Errorf("Failed to write to json file: %w", err)
	}

	return nil
}

func LoadTasks() (string, error) {
	filepath, err := getFilePath()
	if err != nil {
		return "", err
	}
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("Failed to read file: %w", err)
	}
	return string(data), nil
}
