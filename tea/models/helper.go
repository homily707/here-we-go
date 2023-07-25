package models

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type TransferMsg string

func TransferCmd(m string) tea.Cmd {
	return func() tea.Msg {
		return TransferMsg(m)
	}
}

var _ tea.Model = (*ListModel)(nil)

type ListModel struct {
	list.Model
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	model, cmd := m.Model.Update(msg)
	m.Model = model
	return m, cmd
}

// View implements tea.Model.
func (m ListModel) View() string {
	return m.Model.View()
}

type SelectFunc func(Controller, SelectMsg) (Controller, tea.Cmd)

type SubmitFunc func(Controller, FormSubmitMsg) (Controller, tea.Cmd)