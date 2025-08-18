package main

import (
	"fmt"
	"os"
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
		}

	}

}

