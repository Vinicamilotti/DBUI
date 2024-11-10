package components

import (
	interfaces "github.com/Vinicamilotti/DBTUI/interface"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type VimModes uint

const (
	NORMAL VimModes = iota
	INSERT
	VISUAL
)

type QryWritter struct {
	Mode        VimModes
	YetBuffer   string
	Height      int
	Width       int
	CursorLine  int
	CursorIndex int
	Title       string
	TexArea     textarea.Model
	styles      *QryWritterStyle
}
type QryWritterStyle struct {
	BorderColor lipgloss.Color
	Style       lipgloss.Style
}

func BaseStyle() *QryWritterStyle {
	s := new(QryWritterStyle)
	s.BorderColor = lipgloss.Color("36")
	s.Style = lipgloss.NewStyle().
		BorderForeground(s.BorderColor).
		Width(90).
		Height(20).
		BorderStyle(lipgloss.NormalBorder())
	return s
}

func (q *QryWritter) Init() tea.Cmd {
	return nil
}

func (q *QryWritter) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	q.TexArea, cmd = q.TexArea.Update(msg)
	return q, cmd
}

func (q *QryWritter) View() string {
	return lipgloss.JoinVertical(lipgloss.Top, q.Title, q.styles.Style.Render(q.TexArea.View()))
}

func (q *QryWritter) Focus() {
	q.TexArea.Focus()
}
func (q *QryWritter) Value() interface{} {
	return q.TexArea.Value()
}
func CreateQryWritter() interfaces.CustomBubble {
	textarea := textarea.New()
	textarea.SetWidth(90)
	textarea.SetHeight(20)
	return &QryWritter{
		Title:       "Write some Query",
		CursorIndex: 0,
		TexArea:     textarea,
		styles:      BaseStyle(),
	}
}
