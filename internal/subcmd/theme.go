package subcmd

import (
	"fmt"

	"github.com/LordOfTrident/gotheme/internal/state"
	"github.com/LordOfTrident/gotheme/internal/utils"
)

var subcmdTheme = subcmd{
	options: []option{
		{name: "name", req: false},
	},

	desc: func() {
		fmt.Print("Set the current theme to a theme with the given name. If such a\n",
		          "theme does not exist, it will be created.\n")
	},

	run: func(name string, args []string, s *state.State) {
		if len(args) == 0 {
			fmt.Printf("Current theme is %v\n", utils.MakeQuote(s.Current))
			return
		}

		if args[0] == s.Current {
			utils.Die("Theme \"%v\" is already the current theme", s.Current)
		}

		s.Current = args[0]
		if _, ok := s.Themes[s.Current]; !ok {
			for _, ch := range args[0] {
				if !utils.IsNameChar(ch) {
					utils.Die("Theme names only allow letters, spaces, " +
					          "underscores (_) and dashes (-)")
				}
			}

			s.Themes[s.Current] = make(state.Theme)
			utils.Success("Created a new theme \"%v\"", s.Current)
		} else {
			utils.Success("Switched to theme \"%v\"", s.Current)
		}
	},
}
