package main

import (
	"fmt"

	"github.com/estebansalazar-27/notes/files"
	"github.com/estebansalazar-27/notes/notes"
	"github.com/estebansalazar-27/notes/tasks"
)

func main() {
	//  Structs declarations
	var newNote *notes.Note
	var newTask *tasks.Task

	// Get structs data
	title, content, closing := notes.GetNoteData()
	description := tasks.GetTaskData()

	//  Structs definitions
	newTask, newTaskErr := tasks.New(description)
	newNote, newNoteErr := notes.New(title, content, closing)

	if newTaskErr != nil {
		fmt.Print(newTaskErr)
		return
	}
	if newNoteErr != nil {
		fmt.Print(newNoteErr)
		return
	}

	// Saving data
	saveTaskErr := files.OutputData(newNote)
	saveNoteErr := files.OutputData(newTask)
	if saveTaskErr != nil {
		fmt.Print(saveTaskErr)
	}
	if saveNoteErr != nil {
		fmt.Print(saveNoteErr)
	}
	fmt.Print(newNote)
}

func add[T int](a, b T) T {
	return a + b
}
