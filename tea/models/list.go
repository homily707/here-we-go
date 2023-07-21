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

func Test1() Model {
	delegate := list.NewDefaultDelegate()
	delegate.UpdateFunc = selectOne

	v1 := []list.Item{
		item{"a", "a"}, item{"b", "b"}, item{"c", "c"}, item{"d", "d"},
	}
	p0 := list.New(v1, delegate, 0, 0)
	p0.SetShowTitle(true)
	p0.SetShowStatusBar(true)
	p0.SetShowHelp(true)
	p0.Title = "SELECT v1"

	v2 := []list.Item{
		item{"1", "1"}, item{"2", "2"}, item{"3", "3"}, item{"4", "4"},
	}
	p1 := list.New(v2, delegate, 0, 0)
	p1.SetShowTitle(true)
	p1.SetShowStatusBar(true)
	p0.SetShowHelp(true)
	p1.Title = "SELECT v2"

	return Model{
		page0: p0,
		page1: p1,
	}
}

func selectOne(msg tea.Msg, list *list.Model) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			return func() tea.Msg {
				return SelectMsg{
					Item: list.SelectedItem(),
					List: list,
				}
			}
		}
	}

	return nil
}

type SelectMsg struct {
	Item list.Item
	List *list.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case SelectMsg:

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.page0.SetSize(msg.Width, msg.Height)
		m.page1.SetSize(msg.Width, msg.Height)
	}

	m.page0, cmd = m.page0.Update(msg)
	cmds = append(cmds, cmd)
	m.page1, cmd = m.page1.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
}
