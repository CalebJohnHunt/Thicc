package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var space rune = '.'

func (m *model) View() string {
	var gridAndTitle string
	{
		grid := m.drawGrid()
		grid = playingFieldBorderStyle.Render(grid)
		gridAndTitle = lipgloss.JoinVertical(lipgloss.Center, titleStyle.Render("SCROBBLE"), grid)
	}

	var pool string
	{
		pool = letterPoolBorderStyle.Render(formatLetters(m.letters))
	}

	var scores string
	{
		scores = m.getPosScore()
		scores += fmt.Sprintf("\nTotal score: %d", m.score)
		scores = letterPoolBorderStyle.Copy().Width(lipgloss.Width(gridAndTitle) + lipgloss.Width(pool) - 2).Render(scores)
	}

	var controls string
	{
		controls = "CONTROLS"
		controls = lipgloss.JoinVertical(lipgloss.Center, controls, "ctrl+s : swap\nctrl+c : quit")
		controls = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#333333")).
			Width(lipgloss.Width(pool) - 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(c4).
			Render(controls)
	}

	var view string
	{
		view = lipgloss.JoinVertical(lipgloss.Left, pool, controls)
		view = lipgloss.JoinHorizontal(lipgloss.Center, gridAndTitle, view)
		view = lipgloss.JoinVertical(lipgloss.Left, view, scores)
	}

	return view
}

func (m *model) drawGrid() string {

	n := len(m.body)
	_, _ = m.body[n-1], m.colors[n-1]
	sb := strings.Builder{}
	for i, row := range m.body {
		for j := 0; j < len(row); j++ {
			if i == m.y && j == m.x {
				sb.WriteString(m.c.render(string(row[j])))
			} else {
				sb.WriteString(m.colors[i][j].decide().Render(string(row[j])))
			}
			sb.WriteByte(' ')
		}
		if i != n-1 {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

func (m *model) getPosScore() string {
	sb := strings.Builder{}
	cd := m.colors[m.y][m.x]

	sb.WriteRune('↔')
	switch cd.horizontal {
	case SINGLE:
	case NOTWORD:
	case WORD:
		left, right := m.hWord(m.x, m.y)
		word := string(m.body[m.y][left:right])
		score := wordScore(word)
		sb.WriteString(fmt.Sprintf(" %s (%d) ", word, score))
	default:
		log.Fatal("not implemented")
	}
	sb.WriteByte('\n')

	sb.WriteRune('↕')
	switch cd.vertical {
	case SINGLE:
	case NOTWORD:
	case WORD:
		top, bottom := m.vWord(m.x, m.y)
		word := make([]rune, bottom-top)
		for i := top; i < bottom; i++ {
			word[i-top] = m.body[i][m.x]
		}
		score := wordScore(string(word))
		sb.WriteString(fmt.Sprintf(" %s (%d)", string(word), score))
	default:
		log.Fatal("not implemented")
	}

	return sb.String()
}

func formatLetters(letters map[rune]int) string {
	sb := strings.Builder{}
	sb1 := strings.Builder{}

	foo := func(r rune) string {
		sb1.Reset()
		i := letters[r]
		sb1.WriteRune(r)
		sb1.WriteString(": ")
		var n string
		if i > 99 {
			n = "9*"
		} else {
			n = fmt.Sprintf("%02d", i)
		}
		sb1.WriteString(n)
		var txt string
		if i == 0 {
			txt = noLettersLeftStyle.Render(sb1.String())
		} else {
			txt = someLettersLeftStyle.Render(sb1.String())
		}
		return txt
	}

	for r := 'A'; r <= 'M'; r++ {

		sb.WriteString(foo(r))
		sb.WriteString("   ")
		sb.WriteString(foo(r + 13))
		if r != 'M' {
			sb.WriteByte('\n')
		}
	}

	return sb.String()
}
