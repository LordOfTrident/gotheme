package subcmd

import (
	"fmt"

	"github.com/LordOfTrident/gotheme/internal/state"
	"github.com/LordOfTrident/gotheme/internal/utils"
	"github.com/LordOfTrident/gotheme/pkg/table"
)

var subcmdTargets = subcmd{
	desc: func() {
		fmt.Println("List all targets.")
	},

	run: func(name string, args []string, s *state.State) {
		if len(s.Targets) == 0 {
			fmt.Println("There are no targets\n")
			utils.Suggest("You can add a target using add-target, " +
			              "see %v", utils.MakeCmdQuote("-h add-target"))
			return
		}

		var t table.Table
		t.SetTitles("Name", "Path", "Original", "Format")

		for targetName, target := range s.Targets {
			row := []string{
				targetName,
				utils.MakeQuote(target.Path),
				utils.MakeQuote(target.Orig),
				utils.MakeQuote(target.Format),
			}

			t.AddRow(row)
		}
		t.Print()
	},
}
