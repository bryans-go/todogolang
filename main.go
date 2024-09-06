package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Todo struct {
	task   string
	isDone bool
}

var todos []Todo

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n1. Add Todo")
		fmt.Println("2. List Todos")
		fmt.Println("3. Mark Todo as Done")
		fmt.Println("4. Delete Todo")
		fmt.Println("5. Exit")

		fmt.Print("Enter your choice: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addTodo()
		case "2":
			listTodos()
		case "3":
			markTodoDone()
		case "4":
			deleteTodo()
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addTodo() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the task: ")
	scanner.Scan()
	task := scanner.Text()
	todos = append(todos, Todo{task: task, isDone: false})
	fmt.Println("Todo added!")
}

func listTodos() {
	if len(todos) == 0 {
		fmt.Println("No todos to show.")
		return
	}
	fmt.Println("\nYour Todos:")
	for i, todo := range todos {
		status := " "
		if todo.isDone {
			status = "x"
		}
		fmt.Printf("[%d] [%s] %s\n", i+1, status, todo.task)
	}
}

func markTodoDone() {
	scanner := bufio.NewScanner(os.Stdin)
	listTodos()
	fmt.Print("Enter the number of the todo to mark as done: ")
	scanner.Scan()
	input := scanner.Text()
	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(todos) {
		fmt.Println("Invalid number.")
		return
	}
	todos[index-1].isDone = true
	fmt.Println("Todo marked as done!")
}

func deleteTodo() {
	scanner := bufio.NewScanner(os.Stdin)
	listTodos()
	fmt.Print("Enter the number of the todo to delete: ")
	scanner.Scan()
	input := scanner.Text()
	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(todos) {
		fmt.Println("Invalid number.")
		return
	}
	todos = append(todos[:index-1], todos[index:]...)
	fmt.Println("Todo deleted!")
}
