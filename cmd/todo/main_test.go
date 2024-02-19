package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	// Check for windows platform append file extension
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	// Build binary for use
	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run() // Run all test cases below, (functions with *testing.T params)

	fmt.Println("Cleaning up files...")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "test task 1"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	// Combine current path plus binary to run
	// eg: ~/Desktop/todo-cli/todo.exe
	cmdPath := filepath.Join(dir, binName)

	t.Run("AddNewTaskFromArguments", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	task2 := "task 2"
	t.Run("AddNewTaskFromSTDIN", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add")
		// Takes an input from pipe like grep
		// eg: cat file.txt | grep "text"
		cmdStdin, err := cmd.StdinPipe()

		if err != nil {
			t.Fatal(err)
		}

		io.WriteString(cmdStdin, task2)
		cmdStdin.Close()

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		output, err := cmd.CombinedOutput() // Returns the output from the command executed
		if err != nil {
			t.Fatal(err)
		}

		expected := fmt.Sprintf("  1: %s\n  2: %s\n", task, task2)

		if expected != string(output) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(output))
		}
	})

	t.Run("MarkTaskComplete", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-complete", "1")
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DeleteTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-delete", "1")
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

}
