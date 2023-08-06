package subcmd

import (
	"fmt"

	"github.com/LordOfTrident/gotheme/internal/usage"
	"github.com/LordOfTrident/gotheme/internal/utils"
	"github.com/LordOfTrident/gotheme/internal/state"
)

type option struct {
	name string
	req  bool
}

type subcmd struct {
	run     func(string, []string, *state.State)
	options []option
	desc    func()
}

var subcmds = map[string]subcmd{
	"update":        subcmdUpdate,
	"set-color":     subcmdSetColor,
	"remove-color":  subcmdRemoveColor,
	"colors":        subcmdColors,
	"add-target":    subcmdAddTarget,
	"set-target":    subcmdSetTarget,
	"remove-target": subcmdRemoveTarget,
	"targets":       subcmdTargets,
	"theme":         subcmdTheme,
	"remove-theme":  subcmdRemoveTheme,
	"copy-theme":    subcmdCopyTheme,
	"themes":        subcmdThemes,
}

func PrintAll(indent int) {
	for name, _ := range subcmds {
		for i := 0; i < indent; i ++ {
			fmt.Print(" ")
		}

		fmt.Println(name)
	}
}

func PrintUsage(name string) {
	subcmd, ok := subcmds[name]
	if !ok {
		utils.DieTry("-h", "Unknown subcommand \"%v\"", name)
	}

	subcmdUsage := usage.Usage{}
	subcmdUsage  = append(subcmdUsage, name)

	for _, option := range subcmd.options {
		if option.req {
			subcmdUsage = append(subcmdUsage, usage.Req(option.name))
		} else {
			subcmdUsage = append(subcmdUsage, usage.Opt(option.name))
		}
	}

	usage.Gen(usage.Usages{subcmdUsage})
	fmt.Println()
	subcmd.desc()
}

func Run(name string, s *state.State) {
	subcmd, ok := subcmds[name]
	if !ok {
		utils.DieTry("-h", "Unknown subcommand \"%v\"", name)
	}

	args := []string{}
	for _, option := range subcmd.options {
		arg, ok := utils.ShiftArgs()
		if !ok {
			if option.req {
				utils.DieTry("-h " + name, "Expected %v", option.name)
			} else {
				break
			}
		}

		args = append(args, arg)
	}

	if extra, ok := utils.ShiftArgs(); ok {
		utils.DieTry("-h " + name, "Unexpected argument \"%v\"", extra)
	}

	subcmd.run(name, args, s)
}
