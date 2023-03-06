package genfile

import (
	"fmt"
)

func Model(s string) string {
	return fmt.Sprintf(`package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// MainModel is the root state of the app.
type MainModel struct {
	appName string
	spinner spinner.Model
	err error
}

// NewModel configures the initial model at runtime.
func NewModel() MainModel {
	s := spinner.New()
	s.Spinner = spinner.Globe

	return MainModel{
		appName: "%s",
		spinner: s,
	}
}

// Init returns any number of tea.Cmds at runtime.
func (m MainModel) Init() tea.Cmd {
	return m.spinner.Tick
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

	case ErrMsg:
		m.err = msg
		return m, nil

	}

	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

// View renders a string representation of the MainModel.
func (m MainModel) View() string {
	return titleView(m.appName) +
		descView() +
		m.spinner.View() +
		footerView()
}

`, s)
}
