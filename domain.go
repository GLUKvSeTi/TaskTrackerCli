package main

import (
	"time"
)

type Status string

const (
	StatusToDo       Status = "todo"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskStorage struct {
	Tasks  []Task `json:tasks`
	NextID int    `json:"next_id"`
}

type TaskRepository interface {
	Load() (*TaskStorage, error)
	Save(*Task) error
}

type TaskService interface {
	Add(title, description string) error
	Update(id int, title, description string) error
	Delete(id int) error

	UpdateStatus(id int, status Status) error

	//вспомогательные операсьоны
	FindByID(id int) (*Task, error)
	GetAll() []Task
	GetAllByStatus(status Status) []Task
}
