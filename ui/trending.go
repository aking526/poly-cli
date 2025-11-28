package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aking526/poly-cli/polymarket"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Messages
type MarketsLoadedMsg struct {
	Markets []polymarket.Market
}

type ErrMsg struct {
	Err error
}

// TrendingModel handles the trending markets view
type TrendingModel struct {
	client  *polymarket.Client
	markets []polymarket.Market
	loading bool
	err     error
	table   table.Model
}

func NewTrendingModel(client *polymarket.Client) TrendingModel {
	columns := []table.Column{
		{Title: "#", Width: 4},
		{Title: "Question", Width: 60},
		{Title: "Volume", Width: 15},
		{Title: "Price", Width: 10},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(15),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return TrendingModel{
		client:  client,
		loading: true,
		table:   t,
	}
}

func (m TrendingModel) Init() tea.Cmd {
	return fetchTrendingMarkets(m.client)
}

func (m TrendingModel) Update(msg tea.Msg) (TrendingModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case MarketsLoadedMsg:
		m.markets = msg.Markets
		m.loading = false

		rows := []table.Row{}
		for i, market := range m.markets {
			price := "-"
			if market.Price != "" {
				if priceFloat, err := strconv.ParseFloat(market.Price, 64); err == nil {
					price = fmt.Sprintf("%.1f¢", priceFloat*100)
				}
			}

			rows = append(rows, table.Row{
				fmt.Sprintf("%d", i+1),
				market.Question,
				formatVolume(market.Volume),
				price,
			})
		}
		m.table.SetRows(rows)
		return m, nil
	case ErrMsg:
		m.err = msg.Err
		m.loading = false
		return m, nil
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m TrendingModel) View() string {
	var s strings.Builder

	s.WriteString("╭─────────────────────────────────────────────╮\n")
	s.WriteString("│   Polymarket - Trending Markets            │\n")
	s.WriteString("╰─────────────────────────────────────────────╯\n\n")

	if m.loading {
		s.WriteString("  Loading trending markets...\n")
	} else if m.err != nil {
		s.WriteString(fmt.Sprintf("  ❌ Error: %s\n", m.err.Error()))
	} else {
		s.WriteString(baseStyle.Render(m.table.View()) + "\n")
	}

	s.WriteString("─────────────────────────────────────────────\n")
	return s.String()
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

// Helper to refresh manually
func (m TrendingModel) Refresh() (TrendingModel, tea.Cmd) {
	m.loading = true
	m.err = nil
	return m, fetchTrendingMarkets(m.client)
}

const numMarkets = 15

// fetchTrendingMarkets returns a Cmd that fetches trending markets
func fetchTrendingMarkets(client *polymarket.Client) tea.Cmd {
	return func() tea.Msg {
		markets, err := client.GetTrendingMarkets(numMarkets)
		if err != nil {
			return ErrMsg{Err: err}
		}
		return MarketsLoadedMsg{Markets: markets}
	}
}

// formatVolume formats the volume string into a human-readable format
func formatVolume(volumeStr string) string {
	volume, err := strconv.ParseFloat(volumeStr, 64)
	if err != nil {
		return volumeStr
	}

	if volume >= 1_000_000 {
		return fmt.Sprintf("$%.2fM", volume/1_000_000)
	} else if volume >= 1_000 {
		return fmt.Sprintf("$%.2fK", volume/1_000)
	}
	return fmt.Sprintf("$%.2f", volume)
}

