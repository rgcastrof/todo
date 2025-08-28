package tasks

import (
	"todo/storage"
	"encoding/json"
	"fmt"
)

type Manager struct {
	tasks map[int]*Task
}

func NewManager() *Manager {
	m := &Manager{ tasks: make(map[int]*Task) }
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

	m.tasks[newTask.ID] = newTask
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
	if task, ok := m.tasks[taskID]; !task.Done && ok {
		task.Done = true
		return m.save()
	} 
	return fmt.Errorf("Task with ID %d do not found or already done", taskID)
}

func (m *Manager) RmTask(taskID int) error {
	if len(m.tasks) == 0 {
		return fmt.Errorf("No tasks found")
	}
	if task, ok := m.tasks[taskID]; ok {
		fmt.Printf("Task \"%s\" successfully deleted\n", task.Title)
		delete(m.tasks, taskID)
		return m.save()
	}
	return fmt.Errorf("Task with ID %d not found", taskID)
}
