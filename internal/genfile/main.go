package genfile

import "fmt"

func Main(s string) string {
	return fmt.Sprintf(`package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
	tui "%s/internal/tui"
)

func main() {
	m := tui.NewModel()
	p := tea.NewProgram(m)

	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}

`, s)
}
