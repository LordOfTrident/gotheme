package table

import (
	"fmt"
	"strings"
	"regexp"
)

type Table struct {
	longest []int
	titles  []string
	rows    [][]string
}

func (t *Table) SetTitles(titles... string) {
	t.titles = titles

	for i, title := range titles {
		l := len(title)

		if i >= len(t.longest) {
			t.longest = append(t.longest, l)
		} else if l > t.longest[i] {
			t.longest[i] = l
		}
	}
}

func xLen(str string) int {
	r       := regexp.MustCompile("\x1b\\[([^m]*)m")
	matches := r.FindAllStringSubmatch(str, -1)

	l := len(str)
	for _, match := range matches {
		l -= len(match[0])
	}
	return l
}

func (t *Table) AddRow(row []string) {
	t.rows = append(t.rows, row)

	for i, elem := range row {
		l := xLen(elem)

		if i >= len(t.longest) {
			t.longest = append(t.longest, l)
		} else if l > t.longest[i] {
			t.longest[i] = l
		}
	}
}

func (t *Table) Print() {
	fmt.Print("  ")
	for i, title := range t.titles {
		fmt.Printf("\x1b[1;97m%v\x1b[0m", title)

		if i + 1 < len(t.titles) {
			fmt.Printf("%v \x1b[90m|\x1b[0m ", strings.Repeat(" ", t.longest[i] - len(title)))
		}
	}
	fmt.Println()

	for _, row := range t.rows {
		fmt.Print("  ")
		for i, elem := range row {
			fmt.Print(elem)

			if i + 1 < len(t.titles) {
				fmt.Printf("%v \x1b[90m|\x1b[0m ", strings.Repeat(" ", t.longest[i] - xLen(elem)))
			}
		}
		fmt.Println()
	}
}
