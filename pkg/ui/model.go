package ui

import (
	"time"

	"github.com/caarlos0/uhr"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func New(tz string, r *lipgloss.Renderer) (tea.Model, error) {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return nil, err
	}
	indigo := lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	fuschia := lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}
	return model{
		t:       time.Now(),
		tz:      loc,
		indigo:  indigo,
		fuschia: fuschia,
		header:  r.NewStyle().Background(fuschia).Bold(true).Foreground(lipgloss.Color("white")).PaddingRight(1).PaddingLeft(1),
		list:    r.NewStyle().MarginLeft(1),
		italic:  r.NewStyle().Italic(true).Foreground(indigo),
		footer:  r.NewStyle().Foreground(lipgloss.Color("gray")).Faint(true),
	}, err
}

type model struct {
	t                            time.Time
	tz                           *time.Location
	indigo, fuschia              lipgloss.AdaptiveColor
	header, list, italic, footer lipgloss.Style
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
	return m, tea.Every(time.Second, func(t time.Time) tea.Msg {
		return t
	})
}

func (m model) View() string {
	t := m.t.In(m.tz)
	s := m.header.Render("Hallo!") + "\n\n"
	s += "Heute ist " + m.italic.Render(uhr.Weekday(m.t)) + ".\n"
	s += "Es ist jetzt " + m.italic.Render(t.Format("15:04")) + ", aber du kannst auch sagen:\n"
	for _, l := range uhr.Uhr(t) {
		s += m.list.Render("- ") + m.italic.Render(l) + "\n"
	}
	s += "Es ist " + m.italic.Render(uhr.PartOfDay(t)) + ".\n"

	s += m.footer.Render("\ndrücke 'q' zum Beenden")
	return s
}
