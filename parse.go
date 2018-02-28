package cmdr

import (
	"strings"
	"unicode"
)

// Parse parses a string
// resulting into a *Command
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
	lastRune := rune(0)
	cControl := false
	startRune := rune(0)
	f := func(c rune) (r bool) {
		switch {
		case startRune == c:
			startRune = 0
			r = cControl
		case (c == rune('\'') || c == rune('"')) && startRune == 0:
			startRune = c
			cControl = unicode.IsSpace(lastRune)
			r = cControl
		default:
			r = unicode.IsSpace(c) && startRune == 0
		}
		lastRune = c
		return
	}
	return strings.FieldsFunc(cmd, f)
}
