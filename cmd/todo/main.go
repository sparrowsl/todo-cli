package main

import (
	"fmt"
	"os"

	"github.com/sparrowsl/todo-cli"
)

const todoFile = ".todo.json"

func main() {
	list := &todo.List{}

	if err := list.Get(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _, item := range *list {
			fmt.Println(item.Task)
		}

	}

}
