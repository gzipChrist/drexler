package genfile

import "fmt"

func Mod(s string, v string) string {
	return fmt.Sprintf("module %s\n\n%s\n\nrequire (\n\tgithub.com/charmbracelet/bubbles v0.14.0\n\tgithub.com/charmbracelet/bubbletea v0.23.1\n\tgithub.com/charmbracelet/lipgloss v0.6.0\n)\n", s, v)
}
