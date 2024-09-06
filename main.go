package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/bryans-go/todogolang/model"
	"github.com/bryans-go/todogolang/storage"
)

func main() {
	db, err := storage.NewFileStorage("todos.json")
	if err != nil {
		log.Fatalf("Failed to open storage: %v", err)
	}

	m := model.NewModel(db)
	p := tea.NewProgram(m)

	if err := p.Start(); err != nil {
		log.Fatalf("Error starting the program: %v", err)
	}
}
