package genfile

const View = `package tui

import "fmt"

var keyParts = []string{
	" _______",
	"|\\     /",
	"| +---+ ",
	"| |   | ",
	"| |%s  | ",
	"| +---+ ",
	"|/_____\\",
}

func KeySwitchArt(s string) string {
	var art string
	for i, line := range keyParts {
		for _, c := range s {
			if i == 4 {
				art += fmt.Sprintf(line, string(c))
			} else {
				art += line
			}
		}

		if i == 0 {
			art += "\n"
		} else {
			art += "|\n"
		}

	}

	return art
}

`
