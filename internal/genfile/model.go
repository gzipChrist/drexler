package genfile

import (
	"fmt"
)

func Model(s string) string {
	return fmt.Sprintf(`package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// MainModel is the root state of the app.
type MainModel struct {
	appName string
	err error
}

// NewModel configures the initial model at runtime.
func NewModel() MainModel {
	return MainModel{
		appName: "%s",
	}
}

// Init returns any number of tea.Cmds at runtime.
func (m MainModel) Init() tea.Cmd {
	return nil
}

// Update handles all tea.Msgs in the Bubble Tea event loop. 
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	// Handle keypress messages.
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}

	}

	return m, cmd 
}

// View renders a string representation of the MainModel.
func (m MainModel) View() string {
	return MainView(m.appName)  
}

`, s)
}
