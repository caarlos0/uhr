package ui

import (
	"time"

	"github.com/caarlos0/uhr"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func New(tz string) (tea.Model, error) {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return nil, err
	}
	return model{
		t:  time.Now(),
		tz: loc,
	}, err
}

type model struct {
	t  time.Time
	tz *time.Location
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

var (
	indigo  = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	fuschia = lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}
	header  = lipgloss.NewStyle().Background(fuschia).Bold(true).Foreground(lipgloss.Color("white")).PaddingRight(1).PaddingLeft(1)
	list    = lipgloss.NewStyle().MarginLeft(1)
	italic  = lipgloss.NewStyle().Italic(true).Foreground(indigo)
	footer  = lipgloss.NewStyle().Foreground(lipgloss.Color("gray")).Faint(true)
)

func (m model) View() string {
	t := m.t.In(m.tz)
	s := header.Render("Hallo!") + "\n\n"
	s += "Heute ist " + italic.Render(uhr.Weekday(m.t)) + ".\n"
	s += "Es ist jetzt " + italic.Render(t.Format("15.04")) + ", aber du kannst auch sagen:\n"
	for _, l := range uhr.Uhr(t) {
		s += list.Render("- ") + italic.Render(l) + "\n"
	}
	s += "Es ist " + italic.Render(uhr.PartOfDay(t)) + ".\n"

	s += footer.Render("\ndrücke 'q' zum Beenden")
	return s
}
