package cmdr

import (
	"strings"
	"unicode"
)

// var quotes = [...]rune{`'`, `"`}

// Parse parses a string
// resulting into a command
func Parse(cmd string) *Command {

	var (
		cmdStr string
		parts  []string
		args   []string
	)
	parts = parseComplex(cmd)
	cmdStr = parts[0]
	args = parts[1:len(parts)]

	return &Command{
		Command: cmdStr,
		Args:    args,
	}
}

// by @jg_19 on https://t.me/go_br
func parseComplex(cmd string) []string {
	start := false
	startRune := rune(0)
	f := func(c rune) bool {
		switch {
		case c == startRune:
			startRune = 0
			start = !start
			return true
		case (c == rune('\'') || c == rune('"')) && !start:
			startRune = c
			start = !start
			return true
		default:
			return unicode.IsSpace(c) && !start
		}
	}

	return strings.FieldsFunc(cmd, f)
}
