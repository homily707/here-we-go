package kubeutil

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test_strsJoin(t *testing.T) {
	strs := []string{"1", "2", "3", "4", "5"}
	fmt.Println(strsJoin(", ", strs...))
	fmt.Println(strsJoin("\n", strs...))
	fmt.Println(strsJoin("\r", strs...))
	fmt.Println(strsJoin("\t", strs...))
}

func Test_Model(t *testing.T) {
	builder := strings.Builder{}
	for i := 0; i < 100; i++ {
		builder.WriteString(strconv.Itoa(i) + "\n")
	}
	content := builder.String()

	p := tea.NewProgram(
		Model{Screen: ScreenModel{
			Content: content,
		}},
		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
	)

	if err := p.Start(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
