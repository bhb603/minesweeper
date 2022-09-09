package minesweeper

import (
	"bytes"
	"fmt"
	"io"

	tea "github.com/charmbracelet/bubbletea"
)

type CLIGame struct {
	game   *Game
	cursor [2]int
}

func NewCLIGame() *CLIGame {
	return &CLIGame{
		game:   NewGame(Beginner),
		cursor: [2]int{0, 0},
	}
}

func (c *CLIGame) Init() tea.Cmd {
	return nil
}

func (c *CLIGame) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	wrap := func(val, min, max int) int {
		wrapped := val % max
		if wrapped < min {
			wrapped = max + wrapped
		}
		return wrapped
	}

	cmds := []tea.Cmd{}
	var err error

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return c, tea.Quit
		case "up", "k":
			c.cursor[0] = wrap(c.cursor[0]-1, 0, c.game.Height)
		case "down", "j":
			c.cursor[0] = wrap(c.cursor[0]+1, 0, c.game.Height)
		case "left", "h":
			c.cursor[1] = wrap(c.cursor[1]-1, 0, c.game.Width)
		case "right", "l":
			c.cursor[1] = wrap(c.cursor[1]+1, 0, c.game.Width)
		case "enter":
			_, err = c.game.RevealCell(c.cursor[0], c.cursor[1])
		case "r":
			_, err = c.game.RevealAdj(c.cursor[0], c.cursor[1])
		case "f":
			_, err = c.game.ToggleFlagCell(c.cursor[0], c.cursor[1])
		}
	}

	if err != nil {
		cmds = append(cmds, tea.Println("Error: ", err))
	}

	if c.game.Status != GameStatusInProgress {
		cmds = append(cmds, tea.Quit)
	}

	return c, tea.Batch(cmds...)
}

func (c *CLIGame) View() string {
	var buffer bytes.Buffer
	c.printInstructions(&buffer)
	fmt.Fprintln(&buffer, "")
	c.game.PrintHeader(&buffer)
	c.game.PrintGrid(&buffer, c.cursor, false)
	switch c.game.Status {
	case GameStatusLost:
		fmt.Fprintln(&buffer, "you lost")
	case GameStatusWon:
		fmt.Fprintln(&buffer, "you won")
	}

	return buffer.String()
}

func (c *CLIGame) printInstructions(w io.Writer) {
	fmt.Fprintln(w, "Use arrows or hjkl to move the cursor")
	fmt.Fprintln(w, "enter: reveal cell")
	fmt.Fprintln(w, "f:     flag/unflag")
	fmt.Fprintln(w, "r:     reveal adjacent")
	fmt.Fprintln(w, "q:     quit")
}
