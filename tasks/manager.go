package tasks

import (
	"todo/storage"
	"encoding/json"
	"fmt"
)

type Manager struct {
	tasks []Task
}

func NewManager() *Manager {
	return &Manager{ tasks: []Task{} }
}

func (m *Manager) nextID() int {
	lastID := 0
	for _, t := range m.tasks {
		if t.ID > lastID {
			lastID = t.ID
		}
	}
	return lastID + 1
}

func (m *Manager) AddTask(title string) error {

	newTask := &Task{
		ID: m.nextID(),
		Title: title,
		Done: false,
	}

	m.tasks = append(m.tasks, *newTask)

	data, err := json.MarshalIndent(m.tasks, "", "	")
	if err != nil {
		return fmt.Errorf("Failed to serialize tasks: %w", err)
	}
	err = storage.SaveTasks(data)
	if err != nil {
		return fmt.Errorf("Failed to save tasks: %w", err)
	}
	return nil
}
