package utils

import (
	"fmt"
	"os"
	"unicode"
)

var ExecPath, Home string

func init() {
	ExecPath, _ = ShiftArgs()

	Home = os.Getenv("HOME")
	if len(Home) == 0 {
		Die("Failed to get the $HOME environment variable")
	}
}

func Success(format string, args... interface{}) {
	fmt.Printf("\x1b[1;32m[✓] " + format + "\x1b[0m\n", args...)
}

func Suggest(format string, args... interface{}) {
	fmt.Printf("\x1b[1;33m[💡] " + format + "\x1b[0m\n", args...)
}

func Die(format string, args... interface{}) {
	fmt.Fprintf(os.Stderr, "\x1b[1;31m[X] " + format + "\x1b[0m\n", args...)
	os.Exit(1)
}

func DieTry(try string, format string, args... interface{}) {
	Die("%v\n    -> Try %v", fmt.Sprintf(format, args...), MakeCmdQuote(try))
}

func MakeCmdQuote(format string, args... interface{}) string {
	return fmt.Sprintf("\x1b[1;37m\"%v %v\"\x1b[0m", ExecPath, fmt.Sprintf(format, args...))
}

func MakeQuote(format string, args... interface{}) string {
	return fmt.Sprintf("\x1b[1;92m\"" + format + "\"\x1b[0m", args...)
}

func ShiftArgs() (string, bool) {
	if len(os.Args) == 0 {
		return "", false
	}

	arg    := os.Args[0]
	os.Args = os.Args[1:]

	return arg, true
}

func IsNameChar(ch rune) bool {
	return ch == ' ' || ch == '_' || ch == '-' || unicode.IsLetter(ch) || unicode.IsNumber(ch)
}

func YNPrompt(format string, args... interface{}) bool {
	fmt.Printf("\x1b[1;97m%v\x1b[0m (Y/N): ", fmt.Sprintf(format, args...))

	var in string
	fmt.Scanln(&in)

	if len(in) > 0 {
		return in[0] == 'y' || in[0] == 'Y'
	} else {
		return false
	}
}

func ExpandPath(path string) string {
	if path[0] == '~' {
		return Home + path[1:]
	} else {
		return path
	}
}
