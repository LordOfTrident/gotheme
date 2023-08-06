package subcmd

import (
	"fmt"

	"github.com/LordOfTrident/gotheme/internal/state"
)

var subcmdThemes = subcmd{
	desc: func() {
		fmt.Println("List all themes.")
	},

	run: func(name string, args []string, s *state.State) {
		for themeName, _ := range s.Themes {
			if themeName == s.Current {
				fmt.Printf("\x1b[1;97m| %v\x1b[0m\n", themeName)
			} else {
				fmt.Printf("  %v\n", themeName)
			}
		}
	},
}
