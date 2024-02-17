package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/sparrowsl/todo-cli"
)

const todoFile = ".todo.json"

func main() {

	deleteFlag := flag.Int("d", 0, "Delete an item")

	flag.Parse()

	list := todo.List{}

	if err := list.Get(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _, item := range list {
			fmt.Println(item.Task)
		}
	case *deleteFlag > 0:
		if err := list.Delete(*deleteFlag); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		list.Save(todoFile)
	default:
		newItem := strings.Join(os.Args[1:], " ")
		list.Add(newItem)

		if err := list.Save(todoFile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

}
