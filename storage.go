package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileRepository struct {
	fileName string
}

func NewFileRepository(fileName string) *FileRepository {
	return &FileRepository{
		fileName: fileName,
	}
}

func (fr *FileRepository) Load() (*TaskStorage, error) {
	if _, err := os.Stat(fr.fileName); os.IsNotExist(err) {
		return &TaskStorage{
			Tasks:  []Task{},
			NextID: 1,
		}, nil
	}
	data, err := os.ReadFile(fr.fileName)
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения файла %s: %v", fr.fileName, err)
	}
	var store TaskStorage
	if err := json.Unmarshal(data, &store); err != nil {
		return nil, fmt.Errorf("Ошибка декодирования JSON: %v", err)
	}
	return &store, nil
}

func (fr *FileRepository) Save(store *TaskStorage) error {
	data, err := json.MarshalIndent(store, "", " ")
	if err != nil {
		return fmt.Errorf("Ошибка кодирования JSON: %v", err)
	}
	tempFile := fr.fileName + ".tmp"
	if err := os.WriteFile(tempFile, data, 0644); err != nil {
		return fmt.Errorf("Ошибка записи во временный файл: %v", err)
	}
	if err := os.Rename(tempFile, fr.fileName); err != nil {
		os.Remove(tempFile)
		return fmt.Errorf("Ошибка переименования файла: %v", err)
	}
	return nil
}
