package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Model for application state
type model struct{}

func initialModel() model {
  return model{}
}

// Init
func (m model) Init() tea.Cmd {
  return nil
}

// Update
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c", "q":
      return m, tea.Quit
    }
  }
  
  return m, nil
}

// View
func (m model) View() string {
  return "Hello, Polymarket CLI!\n\nPress 'q' to quit.\n"
}

func main() {
  p := tea.NewProgram(initialModel())

  if _, err := p.Run(); err != nil {
    fmt.Println("Error running program:", err)
    os.Exit(1)
  }
}