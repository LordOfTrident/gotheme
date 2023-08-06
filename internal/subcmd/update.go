package subcmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/LordOfTrident/gotheme/internal/state"
	"github.com/LordOfTrident/gotheme/internal/utils"
)

var subcmdUpdate = subcmd{
	desc: func() {
		fmt.Println("Update (process) the target files.")
	},

	run: func(name string, args []string, s *state.State) {
		if len(s.Targets) == 0 {
			fmt.Println("There are no targets to update (process)\n")
			utils.Suggest("You can add a target using add-target, " +
			              "see %v", utils.MakeCmdQuote("-h add-target"))
			return
		}

		for targetName, target := range s.Targets {
			path := utils.ExpandPath(target.Path)
			orig := utils.ExpandPath(target.Orig)

			b, err := os.ReadFile(path)
			if err != nil {
				utils.Die("Failed to open path \"%v\" of target \"%v\"", target.Path, targetName)
			}

			processed := process(string(b), target.Format, s)
			if len(processed) == 0 {
				fmt.Printf("Target \"%v\" has nothing to update\n", targetName)
				continue
			}

			err = os.WriteFile(orig, []byte(processed), 0644)
			if err != nil {
				utils.Die("Failed to write original \"%v\" of target \"%v\"",
				          target.Orig, targetName)
			}

			utils.Success("Updated target \"%v\"", targetName)
		}
	},
}

func process(str, format string, s *state.State) string {
	r       := regexp.MustCompile("\\$\\(([^\\$\\(\\)]*)\\)")
	matches := r.FindAllStringSubmatchIndex(str, -1)

	if len(matches) == 0 {
		return ""
	}

	for i := len(matches) - 1; i >= 0; i -- {
		match := matches[i]

		colorFound := str[match[2]:match[3]]

		color, ok := s.Themes[s.Current][colorFound]
		if !ok {
			continue
		}

		str = str[:match[0]] + color.ToStringFromFormat(format) + str[match[1]:]
	}

	return str
}
