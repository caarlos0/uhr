package main

import (
	"log"

	"github.com/caarlos0/uhr/pkg/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	m, err := ui.New("Local", lipgloss.DefaultRenderer())
	if err != nil {
		log.Fatal(err)
	}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
