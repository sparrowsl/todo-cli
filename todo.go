package todo

import "time"

type item struct {
	Task      string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type List []item

// Marks a to-do item as completed
func (l List) Complete() {}

// Creates a new todo and append it to the list
func (l List) Add() {}

// Deletes a to-do item from the list
func (l List) Delete(id int) {}

// Saves the list of items to a file using the JSON format
func (l List) Save() {}

// Obtains a list of items from a saved JSON file
func (l List) Get() {}
