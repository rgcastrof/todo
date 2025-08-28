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
	m := &Manager{ tasks: []Task{} }
	data, err := storage.LoadTasks()
	if err != nil {
		return m
	}
	err = json.Unmarshal([]byte(data), &m.tasks)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return m
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

func (m *Manager) save() error {
	data, err := json.MarshalIndent(m.tasks, "", "	")
	if err != nil {
		return fmt.Errorf("Failed to serialize tasks: %w", err)
	}
	return storage.SaveTasks(data)
}

func (m *Manager) AddTask(title string) error {

	newTask := &Task{
		ID: m.nextID(),
		Title: title,
		Done: false,
	}

	m.tasks = append(m.tasks, *newTask)
	return m.save()
}

func (m *Manager) ListTasks() {
	if len(m.tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Println("Task List:")
	for _, t := range m.tasks {
		fmt.Printf("ID: %v | Title: %v | Done: %v\n", t.ID, t.Title, t.Done)
	}
}

func (m *Manager) MarkDone(taskID int) error {
	found := false
	for i := range m.tasks {
		if m.tasks[i].ID == taskID && !m.tasks[i].Done {
			m.tasks[i].Done = true
			fmt.Printf("Task \"%s\" marked as done.\n", m.tasks[i].Title)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("Task with ID %d do not found or already done", taskID)
	}

	return m.save()
}

func (m *Manager) RmTask(taskID int) error {
	if len(m.tasks) == 0 {
		return fmt.Errorf("No tasks found")
	}
	for i := range m.tasks {
		if m.tasks[i].ID == taskID {
			taskTitle := m.tasks[i].Title
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			fmt.Printf("Task \"%s\" successfully deleted\n", taskTitle)
			return m.save()
		}
	}
	return fmt.Errorf("Task with ID %d not found", taskID)
}
