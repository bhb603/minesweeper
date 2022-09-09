package main

import (
	"fmt"
	"os"

	"github.com/bhb603/minesweeper/minesweeper"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	prog := tea.NewProgram(minesweeper.NewCLIGame())
	if err := prog.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
