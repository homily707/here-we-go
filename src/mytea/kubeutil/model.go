package kubeutil

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b)
	}()
)

type ModelType int

const (
	INPUTMODE  ModelType = 1
	VISUALMODE ModelType = 2
)

type ScreenModel struct {
	Content string
	//InputText string
	ready      bool
	modelType  ModelType
	viewport   viewport.Model
	inputModel textinput.Model
}

func InitScreenModel() ScreenModel {
	return ScreenModel{
		Content:    "here we go",
		ready:      false,
		modelType:  INPUTMODE,
		viewport:   viewport.Model{},
		inputModel: textinput.New(),
	}
}

func (m ScreenModel) Init() tea.Cmd {
	return nil
}

func (m ScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		inputHeight := lipgloss.Height(m.inputView())
		verticalMarginHeight := headerHeight + footerHeight + inputHeight

		if m.modelType == INPUTMODE {
			m.inputModel.Focus()
		} else if m.modelType == VISUALMODE {
			m.inputModel.Reset()
			m.inputModel.Blur()
		}

		if !m.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.SetContent(m.Content)
			m.ready = true

			// This is only necessary for high performance rendering, which in
			// most cases you won't need.
			//
			// Render the viewport one line below the header.
			m.viewport.YPosition = headerHeight + 1
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}
	case tea.KeyMsg:
		if m.modelType == INPUTMODE {
			switch msg.String() {
			case "ctrl+c":
				// quit
				return m, tea.Quit
			case "esc":
				// change to visual
				m.modelType = VISUALMODE
				m.inputModel.Reset()
				m.inputModel.Blur()
			case "enter":
				m.receiveInput(m.inputModel.Value())
				m.inputModel.Reset()
				m.viewport.GotoBottom()
			}
		} else if m.modelType == VISUALMODE {
			switch msg.String() {
			case "ctrl+c", "q", "esc":
				return m, tea.Quit
			case "i":
				m.modelType = INPUTMODE
			}
		}
	}
	// Handle keyboard and mouse events in the viewport and input
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)
	m.inputModel, cmd = m.inputModel.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *ScreenModel) receiveInput(inputText string) {
	m.Content = m.Content + "\n" + inputText
	m.viewport.SetContent(m.Content)
}

// ===================================================================
// view

func (m ScreenModel) View() string {
	return fmt.Sprintf("%s\n%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView(), m.inputView())
	//return stringsJoin("\n", m.headerView(), m.View(), m.footerView(), m.inputView())
}

func (m ScreenModel) headerView() string {
	title := titleStyle.Render("hello")
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m ScreenModel) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func (m ScreenModel) inputView() string {
	return m.inputModel.View()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func stringsJoin(joiner string, strs ...string) string {
	builder := strings.Builder{}
	for i := 0; i < len(strs); i++ {
		builder.WriteString(strs[i])
		if i != len(strs) {
			builder.WriteString(joiner)
		}
	}
	return builder.String()
}
