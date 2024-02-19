package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sparrowsl/todo-cli"
)

var todoFile = ".todo.json"

func main() {
	listFlag := flag.Bool("list", false, "List all items in todo")
	addFlag := flag.Bool("add", false, "Add new task to the todo")
	completedFlag := flag.Int("complete", 0, "Mark an item as completed")
	deleteFlag := flag.Int("delete", 0, "Delete an item")

	flag.Parse()

	list := &todo.List{}

	if os.Getenv("TODO_FILENAME") != "" {
		todoFile = os.Getenv("TODO_FILENAME")
	}

	if err := list.Get(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *listFlag:
		fmt.Print(list)

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

	case *addFlag:
		task, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		list.Add(task)

		if err := list.Save(todoFile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stderr, "Error: invalid flag!")
		os.Exit(1)
	}

}

// Decides where to get the description of a new task from; args or STDIN
func getTask(reader io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(reader)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	if len(scanner.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}

	return scanner.Text(), nil
}
