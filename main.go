package main

import (
	"fmt"
	"os"

	"github.com/aking526/poly-cli/polymarket"
	"github.com/aking526/poly-cli/ui"
	tea "github.com/charmbracelet/bubbletea"
)

type currentView int

const (
	viewTrending currentView = iota
)

type model struct {
	client       *polymarket.Client
	currentView  currentView
	trendingView ui.TrendingModel
}

func initialModel() model {
	client := polymarket.NewGammaClient()
	return model{
		client:       client,
		currentView:  viewTrending,
		trendingView: ui.NewTrendingModel(client),
	}
}

func (m model) Init() tea.Cmd {
	// Init the current view
	if m.currentView == viewTrending {
		return m.trendingView.Init()
	}
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "r":
			if m.currentView == viewTrending {
				var cmd tea.Cmd
				m.trendingView, cmd = m.trendingView.Refresh()
				return m, cmd
			}
		}
	}

	// Delegate to current view
	if m.currentView == viewTrending {
		m.trendingView, cmd = m.trendingView.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	if m.currentView == viewTrending {
		return m.trendingView.View() + "\nPress 'r' to refresh | 'q' to quit"
	}
	return "Unknown view"
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}