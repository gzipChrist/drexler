package cli

import "errors"

type Command string

const (
	Help   = "help"
	Init   = "init"
	Bubble = "bubble"
)

var m = map[Command]bool{
	Help:   true,
	Init:   true,
	Bubble: true,
}

func (c Command) Validate() error {
	_, ok := m[c]
	if !ok {
		return errors.New("invalid command")
	}

	return nil
}

func (c Command) String() string {
	_, ok := m[c]
	if !ok {
		return ""
	}

	return string(c)
}
