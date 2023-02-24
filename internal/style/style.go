package style

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"time"
)

var header = `

     '||                          '||                  
   .. ||  ... ..    ....  ... ...  ||    ....  ... ..  
 .'  '||   ||' '' .|...||  '|..'   ||  .|...||  ||' '' 
 |.   ||   ||     ||        .|.    ||  ||       ||     
 '|..'||. .||.     '|...' .|  ||. .||.  '|...' .||.
`

var Check = lipgloss.NewStyle().SetString("✓").Foreground(lipgloss.Color("#b7fbd7"))
var X = lipgloss.NewStyle().SetString("𐄂").Foreground(lipgloss.Color("#ff0040"))

func ShowHeader() {
	h := lipgloss.NewStyle().SetString(header).Foreground(lipgloss.Color("#85cbf6"))
	fmt.Printf("%s\n  The TUI Project Generator\n", h)
}

func ShowComplete(t time.Time) {
	fmt.Printf("  %s All done! Generated your TUI project in %v\n", Check, time.Since(t))
}
