package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/estebansalazar-27/notes/user"
)

//	Create a note with the fields
//	Content / title / created at
//
// Get related data from terminal

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Closing   string    `json:"closing"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Save() error {
	//  If u want to export the struc to a file as a json, u must export the fields adding capitalizing the properties
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, marshErr := json.Marshal(note)

	if marshErr != nil {
		return marshErr
	}
	return os.WriteFile(fileName, json, 0644)
}

func New(title, content, closing string) (*Note, error) {
	if title == "" || content == "" || closing == "" {
		return &Note{}, errors.New("All the fields are required to create a valid note, please try again")
	}
	return &Note{
		Title:     title,
		Content:   content,
		Closing:   closing,
		CreatedAt: time.Now(),
	}, nil
}

func GetNoteData() (string, string, string) {
	title := user.GetUserInput("Note's title: ")
	content := user.GetUserInput("Note's content: ")
	closing := user.GetUserInput("Note's closing: ")

	return title, content, closing
}

func (note *Note) Display() {
	fmt.Print(note)
}
