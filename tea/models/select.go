package models

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type choice struct {
	key string   // 唯一确定标识
	name string  // 展示的名字
}

type Choices struct {
	choices []choice
	currentIndex int
	maxIndex int
}

func NewChoices() Choices {
	return Choices{
		maxIndex:     -1,
	}
}

func (c *Choices) AddChoice(k,v string) {
	c.choices = append(c.choices, choice{k, v})
	c.maxIndex ++
}

func (c Choices) Choose() string {
	return c.choices[c.currentIndex].key
}

func (c Choices) Names() []string {
	keys := make([]string, len(c.choices))
	for i, k := range c.choices {
		keys[i] = k.name
	}
	return keys
}

func (c *Choices) Next() bool {
	if c.currentIndex == c.maxIndex {
		return false
	}
	c.currentIndex ++
	return true
}

func (c *Choices) Previous() bool {
	if c.currentIndex == 0 {
		return false
	}
	c.currentIndex --
	return true
}

func StandardString(c Choices) string {
	s := strings.Builder{}
	s.WriteString("Navigate with up/down or k/j. Select with Enter\n\n")

	for i := 0; i < len(c.choices); i++ {
		if c.currentIndex == i {
			s.WriteString("[•] ")
		} else {
			s.WriteString("[ ] ")
		}
		s.WriteString(c.choices[i].name)
		s.WriteString("\n")
	}
	s.WriteString("\n(Press 'q' or 'esc' to quit)\n")

	return s.String()
}

var _ tea.Model = (*SelectModel)(nil)

type SelectModelQuitMsg string

type SelectModel struct {
	choices *Choices
}

func NewSelectModel() SelectModel {
	choices := NewChoices()
	return SelectModel{
		choices: &choices,
	}
}

func (m *SelectModel) AddChoice(k,v string) {
	m.choices.AddChoice(k,v)
}

// Init implements tea.Model.
func (SelectModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m SelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg :
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			choice := m.choices.Choose()
			return m, func() tea.Msg {
				return SelectModelQuitMsg(choice)
			}

		case "down", "j":
			m.choices.Next()

		case "up", "k":
			m.choices.Previous()
		}
	}

	return m, nil
}

// View implements tea.Model.
func (m SelectModel) View() string {
	return StandardString(*m.choices)
}
