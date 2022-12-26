package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	blinkOn  = lipgloss.NewStyle().Background(lipgloss.Color("#ffffff")).Foreground(lipgloss.Color("#000000"))
	blinkOff = lipgloss.NewStyle().Background(lipgloss.Color("#ffffff")).Foreground(lipgloss.Color("#000000"))
	// blinkOff = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))
)

type cursor struct {
	isBlinking bool
	sleep      time.Duration
}

func (c *cursor) render(s string) string {
	b := &blinkOn
	if !c.isBlinking {
		b = &blinkOff
	}
	return b.Render(s)
}

type changeBlink struct{}

func (c *cursor) blink() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(c.sleep)
		c.isBlinking = !c.isBlinking
		return changeBlink{}
	}
}
