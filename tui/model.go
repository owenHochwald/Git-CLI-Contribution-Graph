package tui

import (
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	Actions  list.Model
	Cursor   int
	Selected map[int]struct{}
	Styles   styles
}

const Version = "1.0.0"

func initialModel() Model {
	items := []list.Item{
		item("stats"),
		item("add"),
		item("remove"),
		item("list"),
		item("email"),
	}

	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	m := Model{
		Actions:  l,
		Cursor:   0,
		Selected: make(map[int]struct{}),
	}
	m.updateStyles(true)
	return m
}

func (m *Model) updateStyles(isDark bool) {
	m.Styles = newStyles(isDark)
	m.Actions.Styles.Title = m.Styles.title
	m.Actions.Styles.PaginationStyle = m.Styles.pagination
	m.Actions.Styles.HelpStyle = m.Styles.help
	m.Actions.SetDelegate(itemDelegate{styles: &m.Styles})
}

func (m Model) Init() tea.Cmd {
	// might possibly do an initial command on tui book
	return nil
}
