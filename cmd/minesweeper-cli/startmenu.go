package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type StartMenuModel struct {
	cursor  int
	options []string
	keys    *MenuKeyMap
	help    help.Model
}

func NewStartMenuModel(width int) *StartMenuModel {
	helpModel := help.New()
	helpModel.Width = width
	return &StartMenuModel{
		cursor:  0,
		options: []string{"New Game", "Quit"},
		keys:    NewMenuKeyMap(),
		help:    helpModel,
	}
}

func (m *StartMenuModel) Init() tea.Cmd {
	return nil
}

func (m *StartMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m *StartMenuModel) selectOption() (tea.Model, tea.Cmd) {
	curOption := m.options[m.cursor]
	switch curOption {
	case "New Game":
		model := NewGameMenuModel(m.help.Width)
		return model, nil
	case "Quit":
		return m, tea.Quit
	}

	return m, nil
}

func (m *StartMenuModel) View() string {
	var buffer bytes.Buffer
	m.printList(&buffer)
	printLines(&buffer, 1)
	printHelp(&buffer, m.help, m.keys)
	return buffer.String()
}

func (m *StartMenuModel) printList(w io.Writer) {
	selectedStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("13"))
	for i, item := range m.options {
		if i == m.cursor {
			fmt.Fprintln(w, selectedStyle.Render("> "+item))
		} else {
			fmt.Fprintf(w, "  %s\n", item)
		}
	}
}
