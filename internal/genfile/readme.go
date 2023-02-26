package genfile

import "fmt"

func ReadMe(s string) string {
	return fmt.Sprintf(`# %s

### This project was generated using the [drexler](https://github.com/gzipChrist/drexler) TUI project generator

`, s)
}
