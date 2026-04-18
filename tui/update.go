package tui

import (
	"fmt"
	"git_contribution_cli/utils"

	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < len(m.Actions.Items())-1 {
				m.Cursor++
			}

		// this is where the delegation for each action handler will live
		case "enter", "space":
			_, ok := m.Selected[m.Cursor]
			if ok {
				delete(m.Selected, m.Cursor)
			} else {
				m.Selected[m.Cursor] = struct{}{}
				m.handleAction()
			}
		}

	}

	// no delegate means no update
	return m, nil
}

func HandleListRepos() {
	config, err := utils.LoadConfig()

	if err != nil {
		panic("Something went wrong! Please try again.\n")
	}

	fmt.Printf("Current email: %s\n", config.Email)

	// repos := utils.ParseFileLinesToSlices(utils.GetDotFilePath())
	fmt.Println("Currently tracked repositories:")
	for _, repo := range config.Repos {
		fmt.Printf("  %s\n", repo)
	}
}

func (m Model) handleAction() {
}
