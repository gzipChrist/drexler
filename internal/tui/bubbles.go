package tui

import "errors"

type Bubble string

const (
	TextInput = "textinput"
)

var m = map[Bubble]bool{
	TextInput: true,
}

func (b Bubble) Validate() error {
	_, ok := m[b]
	if !ok {
		return errors.New("invalid bubble")
	}

	return nil
}

func (b Bubble) String() string {
	_, ok := m[b]
	if !ok {
		return ""
	}

	return string(b)
}
