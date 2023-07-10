package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"here-we-go/src/mytea"
)

func main() {
	p := tea.NewProgram(mytea.NewMyModel())
	if err := p.Start(); err != nil {
		fmt.Println("wrong")
	}
}
