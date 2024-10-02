package main

import "github.com/charmbracelet/lipgloss"

// Define UI styles using lipgloss
var (
	titleStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFA500")).Padding(0, 1)
	todoStyle   = lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("#87CEEB"))
	doneStyle   = todoStyle.Copy().Strikethrough(true).Foreground(lipgloss.Color("#00FF00"))
	cursorStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFD700"))
)
