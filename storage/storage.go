package storage

import (
	"encoding/json"
	"os"
	"time"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Done      bool   `json:"done"`
	CreatedAt string `json:"created_at"`
}

func getFilePath() string {
	home, _ := os.UserHomeDir()
	return home + "/.tasks.json"
}

func LoadTasks() ([]Task, error) {
	filePath := getFilePath()

	// If file doesn't exist yet, return an empty list
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(getFilePath(), data, 0644)
}

func NextID(tasks []Task) int {
	max := 0
	for _, t := range tasks {
		if t.ID > max {
			max = t.ID
		}
	}
	return max + 1
}

func NewTask(id int, title string) Task {
	return Task{
		ID:        id,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}