package tasks

import (
	"todo/storage"
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
}
