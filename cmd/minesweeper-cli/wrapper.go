package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type WrapperModel struct {
	inner tea.Model
	keys  *BaseKeyMap
	help  help.Model
}

func NewWrapperModel() *WrapperModel {
	return &WrapperModel{
		inner: NewStartMenuModel(0),
		keys:  NewBaseKeyMap(),
		help:  help.New(),
	}
}

func (m *WrapperModel) Init() tea.Cmd {
	return nil
}

func (m *WrapperModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m.inner, cmd = m.inner.Update(msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}
	}

	return m, cmd
}

func (m *WrapperModel) View() string {
	var buffer bytes.Buffer
	m.printTitle(&buffer)
	fmt.Fprint(&buffer, m.inner.View())
	return buffer.String()
}

func (m *WrapperModel) printTitle(w io.Writer) {
	fmt.Fprintf(w, "Minesweeper\n\n")
}
