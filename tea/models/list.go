package models

import (
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

var _ list.DefaultItem = (*SimpleItem)(nil)

type SimpleItem struct {
	name string
	key  string
}

func (i SimpleItem) Description() string {
	return "key" + i.key
}

// Title implements list.DefaultItem.
func (i SimpleItem) Title() string {
	return i.name
}

func (i SimpleItem) FilterValue() string {
	return i.name
}

func SimpleItemList(names []string) []list.Item {
	items := []list.Item{}
	for _, name := range names {
		items = append(items, SimpleItem{name, name})
	}
	return items
}

func SimpleItemListModel(names []string) ListModel {
	items := SimpleItemList(names)
	list := list.New(items, SelectDelegate(), 0, 0)
	return ListModel{
		Model: list,
	}
}

func SelectDelegate() list.ItemDelegate {
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	delegate.UpdateFunc = selectOne
	return delegate
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


var _ tea.Model = (*Model)(nil)

type Model struct {
	page0, page1 list.Model
	step         int
}

func Test1() Model {
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	delegate.UpdateFunc = selectOne

	v1 := []list.Item{
		SimpleItem{"a", "a"}, SimpleItem{"b", "b"}, SimpleItem{"c", "c"}, SimpleItem{"d", "d"},
	}
	p0 := list.New(v1, delegate, 0, 0)
	p0.SetShowTitle(true)
	p0.SetShowStatusBar(true)
	p0.SetShowHelp(true)
	p0.Title = "SELECT v1"

	v2 := []list.Item{
		SimpleItem{"1", "1"}, SimpleItem{"2", "2"}, SimpleItem{"3", "3"}, SimpleItem{"4", "4"},
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


func (m Model) Init() tea.Cmd {

	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case SelectMsg:
		m.step++

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.page0.SetSize(msg.Width, msg.Height)
		m.page1.SetSize(msg.Width, msg.Height)
	}

	if m.step == 0 {
		m.page0, cmd = m.page0.Update(msg)
		cmds = append(cmds, cmd)
	}

	if m.step == 1 {
		m.page1, cmd = m.page1.Update(msg)
		cmds = append(cmds, cmd)
	}

	if m.step == 2 {
		cmds = append(cmds, 
			tea.ExecProcess(exec.Command("t9k-pf", "nb", "https://proxy.nc201.kube.tensorstack.net/t9k/notebooks/projects/dev-lmh/name/ssh/tree", "3333"), nil))
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.step == 0 {
		return m.page0.View()
	}

	if m.step == 1 {
		return m.page1.View()
	}
	return "end"
}
