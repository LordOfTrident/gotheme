package subcmd

import (
	"fmt"

	"github.com/LordOfTrident/gotheme/internal/state"
	"github.com/LordOfTrident/gotheme/internal/utils"
	"github.com/LordOfTrident/gotheme/pkg/table"
)

var subcmdColors = subcmd{
	desc: func() {
		fmt.Println("List all colors in the current theme.")
	},

	run: func(name string, args []string, s *state.State) {
		if len(s.Themes[s.Current]) == 0 {
			fmt.Println("There are no colors in this theme\n")
			utils.Suggest("You can add a color using set-color, " +
			              "see %v", utils.MakeCmdQuote("-h set-color"))
			return
		}

		var t table.Table
		t.SetTitles("Name", "Hex", "Showcase")

		for colorName, color := range s.Themes[s.Current] {
			row := []string{
				colorName,
				"\x1b[1;92m#" + color.Hex + "\x1b[0m",
				color.Showcase(),
			}

			t.AddRow(row)
		}
		t.Print()
	},
}
