package minesweeper

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keyMap struct {
	Up         key.Binding
	Down       key.Binding
	Left       key.Binding
	Right      key.Binding
	Reveal     key.Binding
	ToggleFlag key.Binding
	Help       key.Binding
	Quit       key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right},
		{
			k.Reveal, k.ToggleFlag,
			k.Help, k.Quit,
		},
	}
}

var defaultKeys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	),
	Reveal: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "reveal cell; if already revealed and all adjacent mines are flagged, reveal adjacent"),
	),
	ToggleFlag: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "flag/unflag"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

type CLIGame struct {
	game   *Game
	cursor [2]int
	help   help.Model
	keys   keyMap
}

func NewCLIGame() *CLIGame {
	return &CLIGame{
		game:   NewGame(Beginner),
		cursor: [2]int{0, 0},
		help:   help.New(),
		keys:   defaultKeys,
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
	case tea.WindowSizeMsg:
		c.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, c.keys.Quit):
			return c, tea.Quit
		case key.Matches(msg, c.keys.Help):
			c.help.ShowAll = !c.help.ShowAll
		case key.Matches(msg, c.keys.Up):
			c.cursor[0] = wrap(c.cursor[0]-1, 0, c.game.Height)
		case key.Matches(msg, c.keys.Down):
			c.cursor[0] = wrap(c.cursor[0]+1, 0, c.game.Height)
		case key.Matches(msg, c.keys.Left):
			c.cursor[1] = wrap(c.cursor[1]-1, 0, c.game.Width)
		case key.Matches(msg, c.keys.Right):
			c.cursor[1] = wrap(c.cursor[1]+1, 0, c.game.Width)
		case key.Matches(msg, c.keys.Reveal):
			var cell *Cell
			cell, err = c.game.GetCell(c.cursor[0], c.cursor[1])
			if err == nil {
				if !cell.Revealed {
					_, err = c.game.RevealCell(c.cursor[0], c.cursor[1])
				} else {
					_, err = c.game.RevealAdj(c.cursor[0], c.cursor[1])
				}
			}
		case key.Matches(msg, c.keys.ToggleFlag):
			_, err = c.game.ToggleFlag(c.cursor[0], c.cursor[1])
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
	fmt.Fprintln(&buffer, "")
	c.game.PrintHeader(&buffer)
	c.game.PrintGrid(&buffer, c.cursor, false)
	c.printHelp(&buffer)
	fmt.Fprintln(&buffer, "")
	switch c.game.Status {
	case GameStatusLost:
		fmt.Fprintln(&buffer, "you lost")
	case GameStatusWon:
		fmt.Fprintln(&buffer, "you won")
	}

	return buffer.String()
}

func (c *CLIGame) printHelp(w io.Writer) {
	fullHelpView := c.help.FullHelpView(c.keys.FullHelp())
	fullHelpHeight := strings.Count(fullHelpView, "\n")

	helpView := c.help.View(c.keys)
	helpHeight := strings.Count(helpView, "\n")

	fmt.Fprint(w, strings.Repeat("\n", fullHelpHeight-helpHeight+1))

	fmt.Fprint(w, helpView)
}
