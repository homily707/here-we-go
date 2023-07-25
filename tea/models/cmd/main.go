package main

import (
	"homily707/here-we-go/tea/models"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// model := models.Test1()

	// model := models.NewForm()
	// model.AddRow("localport", "localport", "")
	// model.AddRow("verbose", "日志等级", "-1")
	// model.AddRow("retryPeriod", "重试周期", "60s")

	namespace := models.SimpleItemListModel("请选择 Project", []string{"dev-lmh", "t9k-system", "demo"})
	resourceType := models.SimpleItemListModel("请选择类型", []string{"Notebook", "Pod", "Service"})
	resourceName := models.SimpleItemListModel("请选择目标", []string{"test", "ssh-test", "deepspeed"})
	
	option := models.NewForm()
	option.AddRow("localport", "localport", "")
	option.AddRow("verbose", "日志等级", "-1")
	option.AddRow("retryPeriod", "重试周期", "60s")

	pf := models.EmptyController()
	pf.AddModel("ns", namespace)
	pf.AddModel("type", resourceType)
	pf.AddModel("name", resourceName)
	pf.AddModel("option", option)

	pf.SetCurrentModel("ns")
	
	pf.WithSelectNext("ns", "type")
	pf.WithSelectNext("type", "name")
	pf.WithSelectNext("name", "option")

	pf.WithSelectFunc("ns", models.StoreItemWhenSelect("namespace"))
	pf.WithSelectFunc("type", models.StoreItemWhenSelect("type"))
	pf.WithSelectFunc("name", models.StoreItemWhenSelect("name"))
	pf.WithSubmitFunc("option", 
		func(c models.Controller, msg models.FormSubmitMsg) (models.Controller, tea.Cmd) {
			port := msg["localport"]
			if port == "" {
				port = "3333"
			}
			execCmd := tea.ExecProcess(exec.Command(
				"t9k-pf", "nb", "https://proxy.nc201.kube.tensorstack.net/t9k/notebooks/projects/dev-lmh/name/ssh/tree", port), nil)
			return c, tea.Batch(tea.ClearScreen, execCmd)
		})
	
	p := tea.NewProgram(pf, tea.WithMouseCellMotion())

	// Run returns the model as a tea.Model.
	p.Run()
}