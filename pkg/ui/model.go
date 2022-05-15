package ui

import (
	"fmt"
	"time"

	"github.com/caarlos0/uhr"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func New() tea.Model {
	return model{
		t: time.Now(),
	}
}

type model struct {
	t time.Time
}

func (m model) Init() tea.Cmd {
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		return t
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case time.Time:
		m.t = msg
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

var (
	bold   = lipgloss.NewStyle().Bold(true)
	list   = lipgloss.NewStyle().Italic(true).MarginLeft(1)
	footer = lipgloss.NewStyle().Foreground(lipgloss.Color("gray")).Faint(true)
)

func (m model) View() string {
	s := bold.Render(fmt.Sprintf(
		"Hallo!\nHeute ist %s.\nEs ist jetzt %s, aber du kannst auch sagen:",
		uhr.Weekday(m.t),
		m.t.Format(time.Kitchen)),
	)
	s += "\n"
	for _, l := range uhr.Uhr(m.t) {
		s += list.Render(fmt.Sprintf("- "+l)) + "\n"
	}

	s += footer.Render("press 'q' to quit")
	return s
}
