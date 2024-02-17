package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
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
	result := m.Run()

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

	// Adding subtests to run
	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, strings.Split(task, " ")...)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath)

		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := task + "\n"

		if expected != string(output) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(output))
		}
	})

}
