package models

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	noStyle      = lipgloss.NewStyle()
)

var _ tea.Model = (*Form)(nil)

type Form struct {
	submitted bool
	status string
	submit string
	focusIndex int
	rowKey []string
	inputMap map[string]textinput.Model
}

type FormSubmitMsg map[string]string

func NewForm() Form {
	status := "use ↑(shift+tab) and ↓(tab) to select, 'enter' to confirm."
	m := make(map[string]textinput.Model)
	return Form{
		status:     status,
		submit:     "[ Submit ]",
		focusIndex: 0,
		rowKey:     []string{},
		inputMap:   m,
	}
}

func (m *Form) AddRow(key, prompt, placeholder string) {
	m.rowKey = append(m.rowKey, key)
	model := textinput.New()
	model.Prompt = prompt + ": "
	model.Placeholder = placeholder
	m.inputMap[key] = model
}

func (m *Form) SetStatus(s string) {
	m.status = s
}

func (m *Form) SetSubmit(s string) {
	m.submit = s
}

func (m Form) Init() tea.Cmd {
	return nil
}

func (m Form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	if m.submitted {
		return m, tea.Quit
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "tab", "down":
			m.focusIndex++
		case "shift+tab", "up":
			m.focusIndex--
		case "enter":
			m. submitted = true
			if m.focusIndex == len(m.rowKey) {
				m.status = "Submitted"
				submitMap := make(map[string]string)
				for k, v := range m.inputMap {
					submitMap[k] = v.Value()
				}
				return m, func() tea.Msg {
					return FormSubmitMsg(submitMap)
				}
			}
			m.focusIndex++
		}
	}

	if m.focusIndex > len(m.rowKey) {
		m.focusIndex = 0
	} else if m.focusIndex < 0 {
		m.focusIndex = len(m.rowKey)
	}

	for i, key := range m.rowKey {
		if i == m.focusIndex {
			input := m.inputMap[key]
			input.Focus()
			input.PromptStyle = focusedStyle
			input, cmd = input.Update(msg)
			cmds = append(cmds, cmd)
			m.inputMap[key] = input
		} else {
			input := m.inputMap[key]
			input.Blur()
			input.PromptStyle = noStyle
			m.inputMap[key] = input
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Form) View() string {
	if m.submitted {
		return ""
	}
	var builder strings.Builder

	builder.WriteString(m.status)
	builder.WriteRune('\n')

	for _, key := range m.rowKey {
		builder.WriteString(m.inputMap[key].View())
		builder.WriteRune('\n')
	}
	botton := m.submit
	if m.focusIndex == len(m.rowKey) {
		botton = focusedStyle.Render(m.submit)
	}
	builder.WriteString(botton)
	return builder.String()
}
