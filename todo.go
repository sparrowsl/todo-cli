package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

// Marks a to-do item as completed
func (l *List) Complete(i int) error {
	ls := *l

	if i <= 0 || i >= len(ls) {
		return fmt.Errorf("Item %d does not exists!!\n", i)
	}

	// ls[i-1] - adjusting the index value for 0 based
	// eg: user selects 1, we select 0 instead, etc
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// Creates a new todo and append it to the list
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

// Deletes a to-do item from the list
func (l *List) Delete(i int) error {
	ls := *l

	if i <= 0 || i >= len(ls) {
		return fmt.Errorf("Item %d does not exists!!\n", i)
	}

	// Delete item in the list
	// adjust the 0 based index again
	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

// Saves the list of items to a file using the JSON format
func (l *List) Save(filename string) error {
	// Marshal the current list items into a JSON of list - json array of item
	jsonMarshal, err := json.Marshal(l)
	if err != nil {
		return err
	}

	// writes to the filename or creates it if file doesnt exists
	return os.WriteFile(filename, jsonMarshal, 0644)
}

// Obtains a list of items from a saved JSON file
func (l *List) Get(filename string) error {
	// Try to read file, return error if file not found
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	// Do nothing if file exists but is empty
	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}
