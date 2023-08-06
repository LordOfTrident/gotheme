package subcmd

import (
	"fmt"

	"github.com/LordOfTrident/gotheme/internal/state"
	"github.com/LordOfTrident/gotheme/internal/utils"
	"github.com/LordOfTrident/gotheme/internal/usage"
	"github.com/LordOfTrident/gotheme/pkg/table"
)

var subcmdAddTarget = subcmd{
	options: []option{
		{name: "name",     req: true},
		{name: "path",     req: true},
		{name: "original", req: true},
		{name: "format",   req: false},
	},

	desc: func() {
		fmt.Print("Add a target with a name, an optional color format, a path and an\n",
		          "original path. The path is a path to the file to be processed, while\n",
		          "the original path is a path to where the processed version of the\n",
		          "file will be put. Default format is " + utils.MakeQuote("#$(X)") + ".\n\n")

		usage.Section("Variables")
		fmt.Println()

		var t table.Table
		t.SetTitles("Format", "Description")

		t.AddRow([]string{"\x1b[1;92m$(X)\x1b[0m", "Hexadecimal color code (RRGGBB)"})
		t.AddRow([]string{"\x1b[1;92m$(R)\x1b[0m", "Decimal 8bit red code"})
		t.AddRow([]string{"\x1b[1;92m$(G)\x1b[0m", "Decimal 8bit green code"})
		t.AddRow([]string{"\x1b[1;92m$(B)\x1b[0m", "Decimal 8bit blue code"})

		t.Print()
	},

	run: func(name string, args []string, s *state.State) {
		for _, ch := range args[0] {
			if !utils.IsNameChar(ch) {
				utils.Die("Target names only allow letters, spaces, underscores (_) and dashes (-)")
			}
		}

		if _, ok := s.Targets[args[0]]; ok {
			utils.Die("A target with the name \"%v\" already exists", args[0])
		}

		target := state.Target{Path: args[1], Orig: args[2]}
		if len(args) == 3 {
			target.Format = "#$(X)"
		} else {
			target.Format = args[3]
		}

		s.Targets[args[0]] = target
		utils.Success("Added new target \"%v\"", args[0])
	},
}
