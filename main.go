package main

import (
	"fmt"
	"os"

	"github.com/bhb603/minesweeper/cli"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	prog := tea.NewProgram(cli.NewWrapperModel())
	if err := prog.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
