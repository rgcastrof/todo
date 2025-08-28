package main

import (
	"fmt"
	"os"
	"strconv"
	"todo/tasks"
)

const (
	version = "0.1.0"
)

func printHelp() {
	fmt.Printf(
		`usage: %s [<args>]

	add  <task description>		- Add a new task
	list				- List all the tasks
	done <id>			- Mark the task as completed
	rm   <id>			- Remove a task
	version				- Shows the program version
	help				- Shows this help message
`, os.Args[0])
}

func handleArgs(m *tasks.Manager) {
	cmd := os.Args[1]
	switch cmd {
	case "add":
		err := m.AddTask(os.Args[2])
		if err != nil {
			fmt.Println(err)
		}
	case "list":
		m.ListTasks()
	case "done":
		taskID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		err = m.MarkDone(taskID)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "del":
		taskID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		err = m.DelTask(taskID)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "version":
		fmt.Printf("%s version: %s\n", os.Args[0], version)
	case "help":
		printHelp()
		return
	}
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
	} else {
		m := tasks.NewManager()
		handleArgs(m)
	}
}
