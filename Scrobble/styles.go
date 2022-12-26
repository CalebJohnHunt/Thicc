package main

import "github.com/charmbracelet/lipgloss"

/*
* f9dbbd
* fca17d
* da627d
* 9a348e
* 0d0628
 */

var (
	c1 = lipgloss.Color("#53DD6C")
	c2 = lipgloss.Color("#63A088")
	c3 = lipgloss.Color("#56638A")
	c4 = lipgloss.Color("#483A58")
	c5 = lipgloss.Color("#B6303D")
)

var (
	/* playing field words */
	goodWordStyle     = lipgloss.NewStyle().Foreground(c1)
	notAWordStyle     = lipgloss.NewStyle().Foreground(c5)
	confusedStyle     = lipgloss.NewStyle().Foreground(c3)
	errorStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffff00"))
	singleLetterStyle = lipgloss.NewStyle()
	/* playing field */
	playingFieldBorderStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(c4)
	/* letter pool */
	noLettersLeftStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#333333"))
	someLettersLeftStyle  = lipgloss.NewStyle()
	letterPoolBorderStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(c4)
	/* title */
	titleStyle = lipgloss.NewStyle().
			Foreground(c2)
)
