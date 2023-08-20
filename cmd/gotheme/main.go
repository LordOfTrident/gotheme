package main

import (
	"fmt"
	"os"

	"github.com/LordOfTrident/gotheme/internal/utils"
	"github.com/LordOfTrident/gotheme/internal/usage"
	"github.com/LordOfTrident/gotheme/internal/subcmd"
	"github.com/LordOfTrident/gotheme/internal/state"
)

const github = "https://github.com/lordoftrident/gotheme"

func help() {
	usage.Gen(usage.Usages{
		{"-h", usage.Opt("subcommand")},
		{usage.Req("subcommand")},
	})
	fmt.Println()

	usage.Section("Github")
	fmt.Printf("\x1b[4;92m%v\x1b[0m\n", github)

	usage.Section("Subcommands")
	fmt.Println()
	subcmd.PrintAll(2)

	fmt.Printf("\n\x1b[1;33mUse %v\x1b[1;33m for help about a specific subcommand\x1b[0m\n",
	           utils.MakeCmdQuote("-h %v", usage.Req("subcommand")))
}

func main() {
	arg, ok := utils.ShiftArgs()
	if !ok {
		utils.DieTry("-h", "Expected a subcommand")
	}

	if arg == "-h" {
		arg, ok := utils.ShiftArgs()

		if !ok {
			help()
		} else {
			if extra, ok := utils.ShiftArgs(); ok {
				utils.DieTry("-h", "Unexpected argument \"%v\"", extra)
			}

			subcmd.PrintUsage(arg)
		}
		os.Exit(0)
	}

	var s state.State
	s.Load()
	subcmd.Run(arg, &s)
	s.Save()
}
