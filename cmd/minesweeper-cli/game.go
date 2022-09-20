package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/bhb603/minesweeper"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type GameModel struct {
	game   *minesweeper.Game
	cursor [2]int
	help   help.Model
	keys   *GameKeyMap
	err    error
}

func NewGameModel(width int, config minesweeper.GameConfig) *GameModel {
	helpModel := help.New()
	helpModel.Width = width
	return &GameModel{
		game:   minesweeper.NewGame(config),
		cursor: [2]int{0, 0},
		help:   helpModel,
		keys:   NewGameKeyMap(),
		err:    nil,
	}
}

func (m *GameModel) Init() tea.Cmd {
	return nil
}

func (m *GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		err error
		cmd tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		if m.game.Status != minesweeper.GameStatusInProgress {
			return NewStartMenuModel(m.help.Width), cmd
		}

		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Up):
			m.cursor[0] = wrapValue(m.cursor[0]-1, m.game.Height)
		case key.Matches(msg, m.keys.Down):
			m.cursor[0] = wrapValue(m.cursor[0]+1, m.game.Height)
		case key.Matches(msg, m.keys.Left):
			m.cursor[1] = wrapValue(m.cursor[1]-1, m.game.Width)
		case key.Matches(msg, m.keys.Right):
			m.cursor[1] = wrapValue(m.cursor[1]+1, m.game.Width)
		case key.Matches(msg, m.keys.Reveal):
			var cell *minesweeper.Cell
			cell, err = m.game.GetCell(m.cursor[0], m.cursor[1])
			if err == nil {
				if !cell.Revealed {
					_, err = m.game.RevealCell(m.cursor[0], m.cursor[1])
				} else {
					_, err = m.game.RevealAdj(m.cursor[0], m.cursor[1])
				}
			}
		case key.Matches(msg, m.keys.ToggleFlag):
			_, err = m.game.ToggleFlag(m.cursor[0], m.cursor[1])
		}

		m.err = err
	}

	return m, cmd
}

func (m *GameModel) View() string {
	var buffer bytes.Buffer

	m.printHeader(&buffer)
	printLines(&buffer, 1)

	m.printGrid(&buffer)
	printLines(&buffer, 1)

	m.printStatus(&buffer)
	printLines(&buffer, 1)

	printHelp(&buffer, m.help, m.keys)

	return buffer.String()
}

func (m *GameModel) printHeader(w io.Writer) {
	height, width := m.game.Height, m.game.Width
	numMines, numFlagged := m.game.NumMines, m.game.NumFlagged
	fmt.Fprintf(w, "%dx%d, %d mines, %d flagged\n", height, width, numMines, numFlagged)
}

func (m *GameModel) printGrid(w io.Writer) {
	height, width := m.game.Height, m.game.Width
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			cell, _ := m.game.GetCell(i, j)
			pre, post, val := " ", " ", " "
			if i == m.cursor[0] && j == m.cursor[1] {
				pre, post = "[", "]"
			}
			if cell.Revealed {
				val = cell.String()
			} else if cell.Flagged {
				val = lipgloss.NewStyle().
					Foreground(lipgloss.Color("13")).
					SetString("⚑").
					String()
			}
			fmt.Fprintf(w, "%s%s%s", pre, val, post)
		}
		fmt.Fprintln(w, "")
	}
}

func (m *GameModel) printStatus(w io.Writer) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("13"))

	if m.err != nil {
		fmt.Fprintln(w, style.Render(m.err.Error()))
	}
	switch m.game.Status {
	case minesweeper.GameStatusLost:
		fmt.Fprintln(w, style.Render("You lost. (Press any key to continue)"))
	case minesweeper.GameStatusWon:
		fmt.Fprintln(w, style.Render("You won. (Press any key to continue)"))
	}
}
