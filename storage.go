package main

import (
    "encoding/json"
    "os"
)

type Todo struct {
    ID    int    `json:"id"`
    Task  string `json:"task"`
    Done  bool   `json:"done"`
}

type Storage interface {
    Load() ([]Todo, error)
    Save(Todo) error
    Delete(int) error
}

type FileStorage struct {
    filename string
}

func NewFileStorage(filename string) (*FileStorage, error) {
    return &FileStorage{filename: filename}, nil
}

func (fs *FileStorage) Load() ([]Todo, error) {
    file, err := os.Open(fs.filename)
    if err != nil {
        if os.IsNotExist(err) {
            return []Todo{}, nil
        }
        return nil, err
    }
    defer file.Close()

    var todos []Todo
    if err := json.NewDecoder(file).Decode(&todos); err != nil {
        return nil, err
    }

    return todos, nil
}

func (fs *FileStorage) Save(t Todo) error {
    todos, err := fs.Load()
    if err != nil {
        return err
    }

    for i, existing := range todos {
        if existing.ID == t.ID {
            todos[i] = t
            return fs.saveAll(todos)
        }
    }

    todos = append(todos, t)
    return fs.saveAll(todos)
}

func (fs *FileStorage) Delete(id int) error {
    todos, err := fs.Load()
    if err != nil {
        return err
    }

    var updated []Todo
    for _, t := range todos {
        if t.ID != id {
            updated = append(updated, t)
        }
    }

    return fs.saveAll(updated)
}

func (fs *FileStorage) saveAll(todos []Todo) error {
    file, err := os.Create(fs.filename)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    return encoder.Encode(todos)
}
