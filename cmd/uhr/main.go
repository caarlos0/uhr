package main

import (
	"log"

	"github.com/caarlos0/uhr/pkg/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m, err := ui.New("Local")
	if err != nil {
		log.Fatal(err)
	}
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
