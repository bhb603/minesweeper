package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// version info is injected into the build via ldflags by default by goreleaser
// `-s -w -X main.version={{.Version}} -X main.commit={{.Commit}} -X main.date={{.Date}} -X main.builtBy=goreleaser`.
var (
	version string = "dev"
)

func main() {
	prog := tea.NewProgram(NewWrapperModel())
	if err := prog.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
