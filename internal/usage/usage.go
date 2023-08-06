package usage

import (
	"fmt"

	"github.com/LordOfTrident/gotheme/internal/utils"
)

type Usage  []string
type Usages []Usage

func Opt(name string) string {
	return "\x1b[35m[" + name + "]\x1b[0m"
}

func Req(name string) string {
	return "\x1b[1;95m<" + name + ">\x1b[0m"
}

func Gen(usages Usages) {
	for i, usage := range usages {
		if i == 0 {
			Section("Usage")
		} else {
			fmt.Printf("       ")
		}

		fmt.Print(utils.ExecPath)
		for _, arg := range usage {
			fmt.Printf(" %v", arg)
		}
		fmt.Println()
	}
}

func Section(title string) {
	fmt.Printf("\x1b[1;36m%v:\x1b[0m ", title)
}
