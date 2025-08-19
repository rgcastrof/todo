package main

import (
	"fmt"
	"os"
	"strconv"
	"todo/tasks"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(os.Args[0])
	} else {
		m := tasks.NewManager()
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
		}

	}

}

