package main

import (
	"homily707/here-we-go/tea/models"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	model := models.NewSelectModel()
	model.AddChoice("nb", "notebook")
	model.AddChoice("pod", "pod")
	model.AddChoice("svc", "service")
	
	p := tea.NewProgram(model)

	// Run returns the model as a tea.Model.
	p.Run()
}