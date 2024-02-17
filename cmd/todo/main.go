package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sparrowsl/todo-cli"
)

const todoFile = ".todo.json"

func main() {
	listFlag := flag.Bool("list", false, "List all items in todo")
	taskFlag := flag.String("task", "", "Add new task to the todo")
	completedFlag := flag.Int("complete", 0, "Mark an item as completed")
	deleteFlag := flag.Int("delete", 0, "Delete an item")

	flag.Parse()

	list := &todo.List{}

	if err := list.Get(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *listFlag:
		for _, item := range *list {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}

	case *deleteFlag > 0:
		if err := list.Delete(*deleteFlag); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		list.Save(todoFile)

	case *completedFlag > 0:
		if err := list.Complete(*completedFlag); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := list.Save(todoFile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *taskFlag != "":
		list.Add(*taskFlag)

		if err := list.Save(todoFile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stderr, "Error: invalid flag!")
		os.Exit(1)
	}

}
