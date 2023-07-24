package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

var _ tea.Model = (*Controller)(nil)

type Controller struct {
	modelMap       map[string]tea.Model
	currentModel   string
	nextWhenSelect map[string]string
	funcWhenSelect map[string]SelectFunc
	funcWhenSubmit map[string]SubmitFunc
	storedValue    map[string]interface{}
}

func EmptyController() Controller {
	return Controller{
		modelMap:       make(map[string]tea.Model),
		currentModel:   "",
		nextWhenSelect: make(map[string]string),
		funcWhenSelect: make(map[string]SelectFunc),
		funcWhenSubmit: make(map[string]SubmitFunc),
	}
}

func (c *Controller) SetCurrentModel(name string) {
	c.currentModel = name
}

func (c *Controller) AddModel(name string, model tea.Model) {
	c.modelMap[name] = model
}

func (c *Controller) StoreValue(key string, value interface{}) {
	c.storedValue[key] = value
}

func (c *Controller) WithSelectNext(name string, next string) {
	c.nextWhenSelect[name] = next
}

func (c *Controller) WithSelectFunc(name string, fn SelectFunc) {
	c.funcWhenSelect[name] = fn
}

func (c *Controller) WithSubmitFunc(name string, fn SubmitFunc) {
	c.funcWhenSubmit[name] = fn
}

func StoreItemWhenSelect(key string) SelectFunc {
	return func(c Controller, msg SelectMsg) (Controller, tea.Cmd) {
		c.StoreValue(key, msg.Item)
		return c, nil
	}
}

func (m Controller) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m Controller) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case SelectMsg:
		if fn, ok := m.funcWhenSelect[m.currentModel]; ok {
			m, cmd = fn(m, msg)
			cmds = append(cmds, cmd)
		}
		// control logic
		if nextModel, ok := m.nextWhenSelect[m.currentModel]; ok {
			cmds = append(cmds, TransferCmd(nextModel))
		}
	case FormSubmitMsg:
		if fn, ok := m.funcWhenSubmit[m.currentModel]; ok {
			m, cmd = fn(m, msg)
			cmds = append(cmds, cmd)
		}
	case tea.WindowSizeMsg:
		for k, v := range m.modelMap {
			if listModel,ok := v.(ListModel); ok {
				listModel.SetSize(msg.Width, msg.Height)
				m.modelMap[k] = listModel
			}
		}
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	curModel, cmd := m.modelMap[m.currentModel].Update(msg)
	m.modelMap[m.currentModel] = curModel
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Controller) View() string {
	return m.modelMap[m.currentModel].View() 
}
