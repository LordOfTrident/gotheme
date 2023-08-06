package subcmd

import (
	"fmt"

	"github.com/LordOfTrident/gotheme/internal/state"
	"github.com/LordOfTrident/gotheme/internal/utils"
)

var subcmdRemoveColor = subcmd{
	options: []option{
		{name: "name", req: true},
	},

	desc: func() {
		fmt.Println("Remove a color with a given name in the current theme.")
	},

	run: func(name string, args []string, s *state.State) {
		if _, ok := s.Themes[s.Current][args[0]]; !ok {
			utils.Die("Color \"%v\" does not exist", args[0])
		}

		if !utils.YNPrompt("Remove color \"%v\"?", args[0]) {
			utils.Die("Cancelled")
		}

		delete(s.Themes[s.Current], args[0])
		utils.Success("Removed color \"%v\"", args[0])
	},
}
