package state

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"encoding/json"

	"github.com/LordOfTrident/gotheme/internal/utils"
)

// TODO: Color aliases

const dataPath = "~/.config/gotheme.json"
var   dataPathFull string

type Color struct {
	Hex     string
	R, G, B uint8
}

func (c *Color) ParseHex(hex string) bool {
	if len(hex) != 6 {
		return false
	}

	color64, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return false
	}

	c.Hex = hex
	c.R   = uint8(color64 & 0xFF0000 >> 16)
	c.G   = uint8(color64 & 0x00FF00 >> 8)
	c.B   = uint8(color64 & 0x0000FF)
	return true
}

func (c *Color) ToStringFromFormat(format string) string {
	str := strings.Replace(format, "$(X)", c.Hex, -1)
	str  = strings.Replace(str,    "$(R)", strconv.Itoa(int(c.R)), -1)
	str  = strings.Replace(str,    "$(G)", strconv.Itoa(int(c.G)), -1)
	str  = strings.Replace(str,    "$(B)", strconv.Itoa(int(c.B)), -1)
	return str
}

func (c *Color) Showcase() string {
	return fmt.Sprintf("\x1b[48;2;%v;%v;%vm  \x1b[0m", c.R, c.G, c.B)
}

type Theme map[string]Color

type Target struct {
	Path, Orig, Format string
}

type State struct {
	Targets map[string]Target
	Themes  map[string]Theme
	Current string
}

func init() {
	dataPathFull = utils.Home + dataPath[1:]
}

func (s *State) Load() {
	b, err := os.ReadFile(dataPathFull)
	if err != nil {
		s.Targets = make(map[string]Target)
		s.Themes  = make(map[string]Theme)
		s.Current = "my-theme"
		s.Themes["my-theme"] = make(Theme)
		return
	}

	err = json.Unmarshal(b, &s)
	if err != nil {
		utils.Die("Failed to parse \"%v\": %v", dataPath, err)
	}
}

func (s *State) Save() {
	b, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		utils.Die("Failed to convert data into JSON")
	}

	err = os.WriteFile(dataPathFull, b, 0644)
	if err != nil {
		utils.Die("Failed to write \"%v\"", dataPath)
	}
}
