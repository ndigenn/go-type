package main

import (
	"fmt"
	"go-type/cmd/parse"
	"go-type/cmd/styles"
	 "math/rand"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)


// passed as context to almost all functions
type model struct{
	width, height int
	textInput textinput.Model
	target string
	tabs []string
	activeTab int
}

// goofy ah functions cause go can't do ternary
func max(a, b int) int {
	if a > b { return a }
	return b
}

func min(a, b int) int {
	if a < b { return a }
	return b
}

// things to be Initialized
func (m *model) Init() tea.Cmd {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 200
	ti.Width = 50

	q := parse.ParseJSON()
	r := rand.Intn(172)
	m.target = q[r].Quote

	m.textInput = ti
	return textinput.Blink
}

// constantly being called, listens for all key strokes
func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "right":
			m.activeTab = min(m.activeTab + 1, len(m.tabs) - 1)
			return m, nil
		case "left":
			m.activeTab = max(m.activeTab - 1, 0)
			return m, nil
		case "ctrl+r":
			m.Init()
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}


// changes the view of tui, changes colors of the target string based on
// what the user typed
func (m model) View() string {
	// Render each tab with active/inactive style
	var renderedTabs []string
	for i, t := range m.tabs {
		style := styles.InactiveTabStyle
		if i == m.activeTab {
			style = styles.ActiveTabStyle
		}
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	// Join tabs horizontally
	tabRow := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	// Content based on active tab
	var content string
	switch m.activeTab {
	case 0: // Typing
		typed := m.textInput.Value()
		var styled string
		for i := 0; i < len(m.target); i++ {
			switch {
			case i < len(typed) && typed[i] == m.target[i]:
				styled += styles.CorrectStyle.Render(string(m.target[i]))
			case i < len(typed) && typed[i] != m.target[i]:
				// if incorrect space typed, put diff character
				char := string(m.target[i])
				if char == " " {
					char = "·"
				}
				styled += styles.IncorrectStyle.Render(char)
			default:
				styled += styles.PendingStyle.Render(string(m.target[i]))
			}
		}
		content = fmt.Sprintf("%s\n\n%s", styled, m.textInput.View())
	case 1: // Stats
		typed := m.textInput.Value()
		correct := 0
		for i := 0; i < len(typed) && i < len(m.target); i++ {
			if typed[i] == m.target[i] {
				correct++
			}
		}
		accuracy := 0.0
		if len(typed) > 0 {
			accuracy = float64(correct) / float64(len(typed)) * 100
		}
		content = fmt.Sprintf("Typed: %d\nCorrect: %d\nAccuracy: %.2f%%", len(typed), correct, accuracy)
	case 2: // About
		content = `
		A TUI typing app built with Bubble Tea and Lip Gloss.
		Author - ndigenn
		Website - ndigenn.com
		github - https://github.com/ndigenn
		`
	}

	// Wrap content in your purple box style
	box := styles.Style.Width(m.width - 4).Render(content)

	paddedTabRow := lipgloss.NewStyle().
	Width(m.width - 4).
	Align(lipgloss.Center).
	Render(tabRow)

	combined := lipgloss.JoinVertical(lipgloss.Left, paddedTabRow, box)

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		combined,
	)
}

func main() {
	tabs := []string{"typing test", "statistics", "info"}
	p := tea.NewProgram(&model{tabs: tabs}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
