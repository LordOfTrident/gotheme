package subcmd

import (
	"fmt"

	"github.com/LordOfTrident/gotheme/internal/utils"
	"github.com/LordOfTrident/gotheme/internal/state"
)

var subcmdSetColor = subcmd{
	options: []option{
		{name: "name", req: true},
		{name: "hex",  req: true},
	},

	desc: func() {
		fmt.Print("Set a color with a given name in the current theme to the given\n",
		          "hex value. Hex string format: \x1b[1;92mRRGGBB\x1b[0m ",
		          "(Example \x1b[92m4ca7b7\x1b[0m).\n")
	},

	run: func(name string, args []string, s *state.State) {
		for _, ch := range args[0] {
			if !utils.IsNameChar(ch) {
				utils.Die("Color names only allow letters, spaces, underscores (_) and dashes (-)")
			}
		}

		color := state.Color{}
		if !color.ParseHex(args[1]) {
			utils.DieTry("-h " + name, "Invalid hex color.")
		}

		if prev, ok := s.Themes[s.Current][args[0]]; ok {
			if !utils.YNPrompt("Overwrite color \"%v\" (\"%v\" %v\x1b[1;97m)?",
			                   args[0], prev.Hex, prev.Showcase()) {
				utils.Die("Cancelled")
			}
		}

		s.Themes[s.Current][args[0]] = color

		utils.Success("Color \"%v\" was set to \"#%v\" %v", args[0], args[1], color.Showcase())
	},
}
