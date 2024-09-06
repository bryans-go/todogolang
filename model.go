package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/bryans-go/todogolang/storage"
)

type Model struct {
	todos []storage.Todo
	db    storage.Storage
}

func NewModel(db storage.Storage) *Model {
	todos, _ := db.Load()
	return &Model{
		todos: todos,
		db:    db,
	}
}

func (m *Model) Init() bubbletea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle user inputs and updates here
	return m, nil
}

func (m *Model) View() string {
	style := lipgloss.NewStyle().Padding(1, 2).Border(lipgloss.NormalBorder())
	return style.Render("Your Todos:\n" + formatTodos(m.todos))
}

func formatTodos(todos []storage.Todo) string {
	// Format the todos into a string
	return ""
}
