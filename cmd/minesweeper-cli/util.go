package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/help"
)

func printHelp(w io.Writer, h help.Model, keys help.KeyMap) {
	fullHelpView := h.FullHelpView(keys.FullHelp())
	fullHelpHeight := strings.Count(fullHelpView, "\n")

	helpView := h.View(keys)
	helpHeight := strings.Count(helpView, "\n")

	printLines(w, fullHelpHeight-helpHeight)

	fmt.Fprint(w, helpView)
}

func printLines(w io.Writer, n int) {
	fmt.Fprint(w, strings.Repeat("\n", n))
}

func wrapValue(val, max int) int {
	wrapped := val % max
	if wrapped < 0 {
		wrapped = max + wrapped
	}
	return wrapped
}
