package genfile

import "fmt"

func ReadMe(s string) string {
	return fmt.Sprintf(`# %s

### This project was generated using the drexler CLI

`, s)
}
