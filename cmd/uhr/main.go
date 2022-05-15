package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/uhr"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

var (
	bold = lipgloss.NewStyle().Bold(true)
	list = lipgloss.NewStyle().Faint(true).Italic(true).MarginLeft(1)
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	t := time.NewTicker(time.Second)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			os.Exit(0)
		case now := <-t.C:
			now = now.Add(-5 * time.Minute)
			termenv.ClearScreen()
			fmt.Println(bold.Render(fmt.Sprintf("Hallo!\nHeute ist %s.\nEs ist...", uhr.Weekday(now))))
			for _, l := range uhr.Uhr(now) {
				fmt.Println(list.Render("- " + l))
			}
		}
	}
}
