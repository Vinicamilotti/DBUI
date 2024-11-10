package interfaces

import tea "github.com/charmbracelet/bubbletea"

type CustomBubble interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() string
	Focus()
	Value() interface{}
}
