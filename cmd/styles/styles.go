package styles

import (
	"github.com/charmbracelet/lipgloss"
)

// styles for actual typed words and target string
var (
	CorrectStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))   // Green
	IncorrectStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))  // Red
	PendingStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))    // Gray

	InactiveTabStyle  = lipgloss.NewStyle().Padding(0, 1).Align(lipgloss.Center)
	ActiveTabStyle = InactiveTabStyle.Foreground(lipgloss.Color("#D70040")).Bold(true)
)

// styles for box
var Style = lipgloss.NewStyle().
	Bold(true).
	// Foreground(lipgloss.Color("#FAFAFA")).
	// Background(lipgloss.Color("#7D56F4")).
	Padding(2).
	PaddingLeft(20).
	PaddingRight(20).
	Border(Border).
	Align(lipgloss.Center).
	Width(50)

// border for box decor
var Border = lipgloss.Border{
	// Top: "-",
	// Bottom: "-",
	 // Left: "|",
	 // Right: "|",
	// TopLeft: "*",
	// TopRight: "*",
	// BottomLeft: "*",
	// BottomRight: "*",
}

func TabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}
