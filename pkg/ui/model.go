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
	bold = lipgloss.NewStyle().Bold(true)
	list = lipgloss.NewStyle().Faint(true).Italic(true).MarginLeft(1)
)

func (m model) View() string {
	s := bold.Render(fmt.Sprintf(
		"Hallo!\nHeute ist %s.\nEs ist jetz %s, aber du kannst auch sagen:\n",
		uhr.Weekday(m.t),
		m.t.Format(time.Kitchen)),
	)
	for _, l := range uhr.Uhr(m.t) {
		s += list.Render(fmt.Sprintf("- %s\n", l))
	}
	return s
}
