package cli

import (
	"bytes"
	"fmt"
	"io"

	"github.com/bhb603/minesweeper/minesweeper"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type GameMenuModel struct {
	cursor  int
	options []string
	keys    *MenuKeyMap
	help    help.Model
}

func NewGameMenuModel(width int) *GameMenuModel {
	helpModel := help.New()
	helpModel.Width = width
	return &GameMenuModel{
		cursor:  0,
		options: []string{"Beginner", "Intermediate", "Expert"},
		keys:    NewMenuKeyMap(),
		help:    helpModel,
	}
}

func (m *GameMenuModel) Init() tea.Cmd {
	return nil
}

func (m *GameMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Up):
			m.cursor = wrapValue(m.cursor-1, len(m.options))
		case key.Matches(msg, m.keys.Down):
			m.cursor = wrapValue(m.cursor+1, len(m.options))
		case key.Matches(msg, m.keys.Select):
			return m.selectOption()
		}
	}

	return m, nil
}

func (m *GameMenuModel) selectOption() (tea.Model, tea.Cmd) {
	curOption := m.options[m.cursor]
	var gameModel tea.Model
	switch curOption {
	case "Beginner":
		gameModel = NewGameModel(m.help.Width, minesweeper.Beginner)
	case "Intermediate":
		gameModel = NewGameModel(m.help.Width, minesweeper.Intermediate)
	case "Expert":
		gameModel = NewGameModel(m.help.Width, minesweeper.Expert)
	}

	return gameModel, nil
}

func (m *GameMenuModel) View() string {
	var buffer bytes.Buffer
	fmt.Fprintln(&buffer, "Select difficulty:")
	m.printList(&buffer)
	printLines(&buffer, 1)
	printHelp(&buffer, m.help, m.keys)
	return buffer.String()
}

func (m *GameMenuModel) printList(w io.Writer) {
	selectedStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("13"))
	for i, item := range m.options {
		if i == m.cursor {
			fmt.Fprintln(w, selectedStyle.Render("> "+item))
		} else {
			fmt.Fprintf(w, "  %s\n", item)
		}
	}
}
