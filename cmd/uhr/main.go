package main

import (
	"log"

	"github.com/caarlos0/uhr/pkg/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(ui.New())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
