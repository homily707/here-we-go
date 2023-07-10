package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

type LineMsg []byte

type StreamViewModel struct {
	viewport viewport.Model
	r        io.Reader
	pgm      *tea.Program
	lines    []string
	ready    bool
}

func NewStreamViewModel(r io.Reader) StreamViewModel {
	return StreamViewModel{r: r}
}

func (m StreamViewModel) Init() tea.Cmd {
	return nil
}

func (m StreamViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case LineMsg:
		m.lines = append(m.lines, string(msg)+"\n")
		m.viewport.SetContent(strings.Join(m.lines, ""))
		m.viewport.GotoBottom()
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return m, tea.Quit
		}
		if k := msg.String(); k == "a" {
			return m, func() tea.Msg {
				return LineMsg([]byte("add\n"))
			}
		}
		if k := msg.String(); k == "b" {
			return m, func() tea.Msg {
				return LineMsg([]byte("b\n"))
			}
		}
	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.HighPerformanceRendering = false
			m.viewport.SetContent("Init")
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
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m StreamViewModel) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

func (m StreamViewModel) headerView() string {
	title := titleStyle.Render("Mr. Pager")
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m StreamViewModel) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// type nonstopReader struct{}

// func (r nonstopReader) Read(p []byte) (n int, err error) {
// 	p = []byte("hello\n")
// 	time.Sleep(2 * time.Second)
// 	return 5, err
// }

func main() {

	r, w := io.Pipe()

	go func() {
		for {
			w.Write([]byte("hello\n"))
			time.Sleep(2 * time.Second)
		}
	}()

	m := NewStreamViewModel(r)
	p := tea.NewProgram(
		m,
		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
		func(p *tea.Program) {
			go func() {
				r := bufio.NewReader(m.r)
				for {
					line, _, err := r.ReadLine()
					if err != nil {
						return
					}
					p.Send(LineMsg(line))
				}
			}()
		},
	)
	m.pgm = p

	if _,err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
