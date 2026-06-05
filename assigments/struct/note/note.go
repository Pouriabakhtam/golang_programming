package note

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func NewNote(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("Title and content cannot be empty")
	}
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil

}

func (n *Note) Output() string {
	return "Title: " + n.Title + "\nContent: " + n.Content + "\nCreated At: " + n.CreatedAt.Format("2006-01-02 15:04:05")
}

func (n *Note) Save() error {
	fileName := strings.ToLower(strings.ReplaceAll(n.Title, " ", "_")) + ".json"
	JsonData, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}
	return os.WriteFile(fileName, JsonData, 0644)
}
