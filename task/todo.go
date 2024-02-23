package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/estebansalazar-27/notes/user"

)

//  Create a note with the fields
//  Content / title / created at
// Get related data from terminal

type Task struct {
	Description string `json:"content"`
}

func (task Task) Save() error {
	//  If u want to export the struc to a file as a json, u must export the fields adding capitalizing the properties
	fileName := "task.json"
	json, marshErr := json.Marshal(task)

	if marshErr != nil {
		return marshErr
	}
	return os.WriteFile(fileName, json, 0644)
}

func New(description string) (*Task, error) {
	if description == "" {
		return &Task{}, errors.New("All the fields are required to create a valid note, please try again")
	}
	return &Task{
		Description: description,
	}, nil
}

func GetTaskData() string {
	description := user.GetUserInput("Task's description")
	return description
}

func (task *Task) Display() {
	fmt.Print(task)
}
