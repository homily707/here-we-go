package models

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type item struct {
	name string
	key  string
}

func (i item) FilterValue() string {
	return i.name
}

var _ tea.Model = (*Model)(nil)

type Model struct {
	page0, page1 list.Model
}

func New() Model {
	delegate := list.NewDefaultDelegate()
	listModel := list.New([]list.Item{}, delegate, 0, 0)
	listModel.SetShowTitle(true)
	listModel.SetShowStatusBar(true)
	listModel.Title = "SELECT"

}

func selectOne(msg tea.Msg, list *list.Model) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			return
		}
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

}

func (m Model) View() string {
}
