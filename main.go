package main

import (
	"database/sql"

	"github.com/Vinicamilotti/DBTUI/components"
	"github.com/Vinicamilotti/DBTUI/db"
	interfaces "github.com/Vinicamilotti/DBTUI/interface"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type DBTUI struct {
	Width          int
	Height         int
	Database       *sql.DB
	QryWritter     interfaces.CustomBubble
	TableManager   interfaces.CustomBubble
	ComponentIndex int
}

func (p *DBTUI) HandleQueryComponent(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "f5":
			sql := p.QryWritter.Value().(string)
			test, err := p.Database.Query(sql)
			if err != nil {
				panic(err)
			}
			p.TableManager.(*components.TableManager).SetDataset("test", test)
			tModel, tcmd := p.TableManager.Update(msg)
			p.TableManager = tModel.(interfaces.CustomBubble)
			p.ComponentIndex = 1
			return p, tcmd
		}
	}

	p.QryWritter.Focus()
	teaModel, cmd := p.QryWritter.Update(msg)
	p.QryWritter = teaModel.(interfaces.CustomBubble)
	return p, cmd

}

func (p *DBTUI) Init() tea.Cmd {
	return nil
}

func (p *DBTUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.Height = msg.Height
		p.Width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return p, tea.Quit
		}
	}
	switch p.ComponentIndex {
	case 0:
		return p.HandleQueryComponent(msg)
	case 1:
		p.TableManager.(*components.TableManager).Table.Focus()
	}

	return p, nil
}

func (p *DBTUI) View() string {
	qryWritter := lipgloss.JoinVertical(lipgloss.Center, p.QryWritter.View())
	tableManager := p.TableManager.View()
	mainPanel := lipgloss.JoinVertical(lipgloss.Center, qryWritter, tableManager)
	return lipgloss.Place(p.Width,
		p.Height,
		lipgloss.Center,
		lipgloss.Top,
		mainPanel)
}

func InitProgram() *DBTUI {
	return &DBTUI{
		QryWritter:     components.CreateQryWritter(),
		TableManager:   components.CreateTableManager(),
		Database:       db.CreateConnection(),
		ComponentIndex: 0,
	}
}
func main() {
	p := tea.NewProgram(InitProgram(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
