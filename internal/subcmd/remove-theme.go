package subcmd

import (
	"fmt"

	"github.com/LordOfTrident/gotheme/internal/state"
	"github.com/LordOfTrident/gotheme/internal/utils"
)

var subcmdRemoveTheme = subcmd{
	options: []option{
		{name: "name", req: true},
	},

	desc: func() {
		fmt.Println("Remove the given theme. It must not be the currently used theme.")
	},

	run: func(name string, args []string, s *state.State) {
		if args[0] == s.Current {
			utils.Die("Cannot remove the current theme")
		}

		if _, ok := s.Themes[args[0]]; !ok {
			utils.Die("Theme \"%v\" does not exist", args[0])
		}

		if !utils.YNPrompt("Remove theme \"%v\"?", args[0]) {
			utils.Die("Cancelled")
		}

		delete(s.Themes, args[0])
		utils.Success("Removed theme \"%v\"", args[0])
	},
}
