package subcmd

import (
	"fmt"

	"github.com/LordOfTrident/gotheme/internal/state"
	"github.com/LordOfTrident/gotheme/internal/utils"
)

var subcmdCopyTheme = subcmd{
	options: []option{
		{name: "name", req: true},
	},

	desc: func() {
		fmt.Println("Copy the current theme with a new name.")
	},

	run: func(name string, args []string, s *state.State) {
		if _, ok := s.Themes[args[0]]; ok {
			utils.Die("Theme \"%v\" already exists", args[0])
		}

		s.Themes[args[0]] = make(state.Theme)
		for colorName, color := range s.Themes[s.Current] {
			s.Themes[args[0]][colorName] = color
		}
		utils.Success("Copied current theme \"%v\" as \"%v\"", s.Current, args[0])
	},
}
