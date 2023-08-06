package subcmd

import (
	"fmt"

	"github.com/LordOfTrident/gotheme/internal/utils"
	"github.com/LordOfTrident/gotheme/internal/state"
	"github.com/LordOfTrident/gotheme/internal/usage"
)

var subcmdSetTarget = subcmd{
	options: []option{
		{name: "name",  req: true},
		{name: "key",   req: true},
		{name: "value", req: true},
	},

	desc: func() {
		fmt.Println("Set the value of a key of a given target.\n")
		usage.Section("Keys")
		fmt.Println("  path\n  original\n  format")
	},

	run: func(name string, args []string, s *state.State) {
		target, ok := s.Targets[args[0]]
		if !ok {
			utils.Die("Target \"%v\" does not exist")
		}

		switch args[1] {
		case "path":     target.Path   = args[2]
		case "original": target.Orig   = args[2]
		case "format":   target.Format = args[2]

		default: utils.DieTry("-h " + name, "Unknown key \"%v\"", args[1])
		}

		s.Targets[args[0]] = target
		utils.Success("Key \"%v\" of target \"%v\" was set to \"%v\"", args[1], args[0], args[2])
	},
}
