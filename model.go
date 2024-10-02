package main

import (
	"fmt"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh/forms"
	"github.com/charmbracelet/huh/lists"
)

type Todo struct {
	task   string
	isDone bool
}

var todos []Todo

// Model struct representing the state of the UI
type model struct {
	choice    int
	inputMode bool
	form      forms.Model
	list      lists.Model
	cursor    int
}

// Initial model
func initialModel() model {
	// Create a new form for adding TODOs
	form := forms.New()
	form.Input("task", "Enter your task")

	// Create an initial empty list
	list := lists.New()

	return model{
		choice:    0,
		inputMode: false,
		form:      form,
		list:      list,
		cursor:    0,
	}
}

// Init method for the tea.Model interface
func (m model) Init() tea.Cmd {
	// No initial command
	return nil
}

// Update method to handle key inputs and update the model state
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Quit the program
		case "q", "ctrl+c":
			return m, tea.Quit

		// Navigate the menu with up/down arrow keys
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < 4 {
				m.cursor++
			}

		// Handle menu selection with Enter key
		case "enter":
			switch m.cursor {
			case 0: // Add Todo
				m.inputMode = true
			case 1: // List Todos
				m.updateListView()
			case 2: // Mark Todo as Done
				m.markTodoAsDone()
			case 3: // Delete Todo
				m.deleteTodo()
			case 4: // Exit
				return m, tea.Quit
			}

		// Handle form input for adding a new TODO task
		default:
			if m.inputMode {
				var doneCmd tea.Cmd
				m.form, doneCmd = m.form.Update(msg)
				if doneCmd != nil {
					m.addTodo()
					m.inputMode = false
				}
			}
		}
	}

	return m, nil
}

// View method renders the UI
func (m model) View() string {
	if m.inputMode {
		// Render form if in input mode
		return m.form.View()
	}

	// Render menu options
	menu := []string{"Add Todo", "List Todos", "Mark Todo as Done", "Delete Todo", "Exit"}
	var menuView string
	for i, option := range menu {
		cursor := " " // No cursor by default
		if m.cursor == i {
			cursor = ">"
		}
		menuView += fmt.Sprintf("%s %s\n", cursor, option)
	}

	if len(todos) > 0 {
		// Render TODO list if available
		menuView += "\n" + m.list.View()
	} else {
		menuView += "\nNo todos to show."
	}

	return menuView
}

// Add a new todo to the list
func (m *model) addTodo() {
	task := m.form.Values()["task"]
	if task != "" {
		todos = append(todos, Todo{task: task, isDone: false})
		m.form.Clear()
		m.updateListView()
	}
}

// Update the list view with todos
func (m *model) updateListView() {
	items := make([]lists.Item, len(todos))
	for i, todo := range todos {
		status := "[ ]"
		if todo.isDone {
			status = "[x]"
		}
		items[i] = lists.NewItem(fmt.Sprintf("%d. %s %s", i+1, status, todo.task))
	}
	m.list.SetItems(items)
}

// Mark a todo as done
func (m *model) markTodoAsDone() {
	var input string
	fmt.Print("Enter the number of the todo to mark as done: ")
	fmt.Scanln(&input)
	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(todos) {
		fmt.Println("Invalid number.")
		return
	}
	todos[index-1].isDone = true
	m.updateListView()
	fmt.Println("Todo marked as done!")
}

// Delete a todo
func (m *model) deleteTodo() {
	var input string
	fmt.Print("Enter the number of the todo to delete: ")
	fmt.Scanln(&input)
	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(todos) {
		fmt.Println("Invalid number.")
		return
	}
	todos = append(todos[:index-1], todos[index:]...)
	m.updateListView()
	fmt.Println("Todo deleted!")
}
