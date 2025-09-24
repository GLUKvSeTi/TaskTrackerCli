package main

import (
	"fmt"
	"time"
)

type TaskServiceImpl struct {
	repository TaskRepository
	store      *TaskStorage
}

func NewTaskService(repository TaskRepository) (TaskService, error) {
	service := &TaskServiceImpl{
		repository: repository,
	}
	store, err := repository.Load()
	if err != nil {
		return nil, fmt.Errorf("Ошибка инициализации сервиса: %v", err)
	}

	service.store = store

	return service, nil
}

func (ts *TaskServiceImpl) save() error {
	if err := ts.repository.Save(ts.store); err != nil {
		return fmt.Errorf("Ошибка сохранения данных: %v", err)
	}
	return nil
}

func (ts *TaskServiceImpl) Add(title, description string) error {
	task := Task{
		ID:          ts.store.NextID,
		Title:       title,
		Description: description,
		Status:      StatusToDo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	ts.store.Tasks = append(ts.store.Tasks, task)
	ts.store.NextID++

	return nil
}

func (ts *TaskServiceImpl) Update(id int, title, description string) error {
	for i := range ts.store.Tasks {
		if ts.store.Tasks[i].ID == id {
			ts.store.Tasks[i].Title = title
			ts.store.Tasks[i].Description = description
			ts.store.Tasks[i].UpdatedAt = time.Now()
			return ts.save()
		}
	}
	return fmt.Errorf("404 not found: %d", id)
}

func (ts *TaskServiceImpl) Delete(id int) error {
	for i, task := range ts.store.Tasks {
		if task.ID == id {
			ts.store.Tasks = append(ts.store.Tasks[:i], ts.store.Tasks[i+1:]...)
			return ts.save()
		}
	}

	return fmt.Errorf("404 not found: %d", id)
}
