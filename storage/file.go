package storage

import (
	"fmt"
	"os"
)

func SaveTasks(data []byte) error {
	file, err := os.Create("tasks.json")
	if err != nil {
		return fmt.Errorf("Failed to create json file: %w", err)
	}
	defer file.Close()

	return os.WriteFile("tasks.json", data, 0644)
}
