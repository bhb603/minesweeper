package cli

import (
	"github.com/charmbracelet/bubbles/key"
)

var (
	QuitKeyBinding = key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	)
	HelpKeyBinding = key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	)
	UpKeyBinding = key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	)
	DownKeyBinding = key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	)
	LeftKeyBinding = key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	)
	RightKeyBinding = key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	)
	SelectKeyBinding = key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	)
	RevealKeyBinding = key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "reveal cell OR reveal adjacent if all adjacent mines are flagged"),
	)
	ToggleFlagKeyBinding = key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "flag/unflag"),
	)
)

type BaseKeyMap struct {
	Quit key.Binding
	Help key.Binding
}

func NewBaseKeyMap() *BaseKeyMap {
	return &BaseKeyMap{
		Quit: QuitKeyBinding,
		Help: HelpKeyBinding,
	}
}

func (km *BaseKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{km.Quit, km.Help}
}

func (km *BaseKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{km.Quit, km.Help},
		{},
	}
}

type MenuKeyMap struct {
	*BaseKeyMap
	Up     key.Binding
	Down   key.Binding
	Select key.Binding
}

func NewMenuKeyMap() *MenuKeyMap {
	return &MenuKeyMap{
		BaseKeyMap: NewBaseKeyMap(),
		Up:         UpKeyBinding,
		Down:       DownKeyBinding,
		Select:     SelectKeyBinding,
	}
}

func (km *MenuKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{km.Up, km.Down, km.Select, km.Quit, km.Help},
		{},
	}
}

type GameKeyMap struct {
	*BaseKeyMap
	Up         key.Binding
	Down       key.Binding
	Left       key.Binding
	Right      key.Binding
	Reveal     key.Binding
	ToggleFlag key.Binding
}

func NewGameKeyMap() *GameKeyMap {
	return &GameKeyMap{
		BaseKeyMap: NewBaseKeyMap(),
		Up:         UpKeyBinding,
		Down:       DownKeyBinding,
		Left:       LeftKeyBinding,
		Right:      RightKeyBinding,
		Reveal:     RevealKeyBinding,
		ToggleFlag: ToggleFlagKeyBinding,
	}
}

func (km *GameKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{km.Up, km.Down, km.Left, km.Right},
		{km.Reveal, km.ToggleFlag, km.Quit, km.Help},
	}
}
