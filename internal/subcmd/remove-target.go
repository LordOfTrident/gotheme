package subcmd

import (
	"fmt"

	"github.com/LordOfTrident/gotheme/internal/state"
	"github.com/LordOfTrident/gotheme/internal/utils"
)

var subcmdRemoveTarget = subcmd{
	options: []option{
		{name: "name", req: true},
	},

	desc: func() {
		fmt.Println("Remove a target.")
	},

	run: func(name string, args []string, s *state.State) {
		if _, ok := s.Targets[args[0]]; !ok {
			utils.Die("Target \"%v\" does not exist", args[0])
		}

		if !utils.YNPrompt("Remove target \"%v\"?", args[0]) {
			utils.Die("Cancelled")
		}

		delete(s.Targets, args[0])
		utils.Success("Removed target \"%v\"", args[0])
	},
}
