package main

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Dir int

const (
	LEFT  = Dir(-1)
	RIGHT = Dir(1)
	UP    = Dir(-2)
	DOWN  = Dir(2)
)

type wordType int

const (
	SINGLE = wordType(iota)
	WORD
	NOTWORD
)

type colorDecider struct {
	vertical, horizontal wordType
}

func (c *colorDecider) decide() *lipgloss.Style {
	switch {
	case c.vertical == SINGLE && c.horizontal == SINGLE:
		return &singleLetterStyle
	case (c.vertical == WORD && c.horizontal != NOTWORD) ||
		(c.horizontal == WORD && c.vertical != NOTWORD):
		return &goodWordStyle
	case (c.vertical == NOTWORD && c.horizontal != WORD) ||
		(c.horizontal == NOTWORD && c.vertical != WORD):
		return &notAWordStyle
	case (c.vertical == WORD && c.horizontal == NOTWORD) ||
		(c.horizontal == WORD && c.vertical == NOTWORD):
		return &confusedStyle
	default:
		log.Println("Error color")
		return &errorStyle
	}
}

type model struct {
	body       [][]rune
	c          cursor
	colors     [][]colorDecider
	debounce   *time.Timer
	dictionary map[string]struct{}
	dir        Dir
	letters    map[rune]int
	memo       string
	score      int
	x, y       int
}

func (m *model) Init() tea.Cmd {
	return m.c.blink()
}

func (m *model) resetDebounce() {
	if m.debounce != nil {
		m.debounce.Stop()
	}
	m.debounce = time.AfterFunc(time.Millisecond*100, func() {
		m.score = m.calculateScore()
	})
}

func (m *model) reset() {
	for i := range m.body {
		for j := range m.body[i] {
			m.body[i][j] = space
			m.colors[i][j].horizontal = SINGLE
			m.colors[i][j].vertical = SINGLE
		}
	}
	m.x = 0
	m.y = 0
	m.dir = RIGHT
}

func (m *model) move() {
	m.moveDir(m.dir)
}

func (m *model) moveDir(dir Dir) {
	switch {
	case dir == UP && m.y > 0:
		m.y--
	case dir == LEFT && m.x > 0:
		m.x--
	case dir == DOWN && m.y < len(m.body)-1:
		m.y++
	case dir == RIGHT && m.x < len(m.body[m.y])-1:
		m.x++
	}
}

func (m *model) hWord(x, y int) (start, end int) {
	start, end = x-1, x+1
	// go left
	for ; start >= 0 && m.body[y][start] != space; start-- {
	}
	start++
	if start < 0 {
		start = 0
	}
	// go right
	for ; end < len(m.body[y]) && m.body[y][end] != space; end++ {
	}
	if end > len(m.body[y]) {
		end = len(m.body[y])
	}
	return start, end
}

func (m *model) vWord(x, y int) (start, end int) {
	start, end = y-1, y+1
	// go up
	for ; start >= 0 && m.body[start][x] != space; start-- {
	}
	start++
	if start < 0 {
		start = 0
	}
	// go down
	for ; end < len(m.body) && m.body[end][x] != space; end++ {
	}
	if end > len(m.body) {
		end = len(m.body)
	}
	return start, end
}

func (m *model) calculateScore() (score int) {
	for y, row := range m.colors {
		for x, cd := range row {
			if cd.horizontal == WORD {
				score += LETTER_SCORES[m.body[y][x]]
			}
			if cd.vertical == WORD {
				score += LETTER_SCORES[m.body[y][x]]
			}
		}
	}
	return score
}

func (m *model) determineWordType(word string) wordType {
	if len(word) <= 1 {
		return SINGLE
	}

	if _, ok := m.dictionary[word]; ok {
		return WORD
	}

	return NOTWORD
}

// We know that m.body[y][x] != space
func (m *model) foo(x, y int) {
	// row
	{
		left, right := m.hWord(x, y)

		t := m.determineWordType(string(m.body[y][left:right]))

		for l := left; l < right; l++ {
			m.colors[y][l].horizontal = t
		}
	}

	// col
	{
		top, bottom := m.vWord(x, y)

		word := make([]rune, bottom-top)
		for l := top; l < bottom; l++ {
			word[l-top] = m.body[l][x]
		}

		t := m.determineWordType(string(word))

		for l := top; l < bottom; l++ {
			m.colors[l][x].vertical = t
		}
	}
}

func (m *model) updateCross(x, y int) {
	if m.body[y][x] != space {
		m.foo(x, y)
	} else {
		m.colors[y][x].horizontal = SINGLE
		m.colors[y][x].vertical = SINGLE
		if x > 0 && m.body[y][x-1] != space {
			m.foo(x-1, y)
		}
		if x < len(m.body[y])-1 && m.body[y][x+1] != space {
			m.foo(x+1, y)
		}
		if y > 0 && m.body[y-1][x] != space {
			m.foo(x, y-1)
		}
		if y < len(m.body)-1 && m.body[y+1][x] != space {
			m.foo(x, y+1)
		}
	}
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+r":
			m.reset()
		case "ctrl+j":
			m.memo = "test"
		case " ":
			m.move()
		// case "ctrl+s":
		// 	m.memo = strconv.Itoa(m.calculateScore())
		case "backspace":
			if m.body[m.y][m.x] == space {
				m.moveDir(-m.dir)
				if m.body[m.y][m.x] != space {
					m.letters[m.body[m.y][m.x]]++
				}
				m.body[m.y][m.x] = space
				m.updateCross(m.x, m.y)
			} else {
				m.letters[m.body[m.y][m.x]]++
				m.body[m.y][m.x] = space
				m.updateCross(m.x, m.y)
				m.moveDir(-m.dir)
			}
			m.resetDebounce()
		case "delete":
			if m.body[m.y][m.x] != space {
				m.letters[m.body[m.y][m.x]]++
				m.body[m.y][m.x] = space
				m.updateCross(m.x, m.y)
				m.resetDebounce()
			}
		case "right":
			m.dir = RIGHT
			m.moveDir(RIGHT)
		case "left":
			m.dir = RIGHT
			m.moveDir(LEFT)
		case "up":
			m.dir = DOWN
			m.moveDir(UP)
		case "down":
			m.dir = DOWN
			m.moveDir(DOWN)
		case "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
			"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z":
			// "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "`", "~", ",", "<", ".", ">", "/", "?", ";", ":", "'",
			// `"`, "[", "{", "]", "}", `\`, "|", "-", "_", "=", "+", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")"

			r := rune(msg.String()[0])

			// to upper
			if 'a' <= r && r <= 'z' {
				r += 'A' - 'a'
			}

			if m.memo == "test" && m.letters[r] > 0 {
				m.letters[r]--
				for i := 0; i < 3; i++ {
					m.letters[rune(rs.Int31n(26)+'A')]++
				}
				break
			}

			if m.letters[r] > 0 {
				m.letters[r]--
				// Replacing
				if m.body[m.y][m.x] != space {
					m.letters[m.body[m.y][m.x]]++
				}
				m.body[m.y][m.x] = r
				m.updateCross(m.x, m.y)
				m.move()
			}
			m.resetDebounce()
		default:
			log.Println("New key: ", msg.String())
		}
	case changeBlink:
		cmds = tea.Batch(cmds, m.c.blink())
	default:
		log.Printf("New message: %T %v", msg, msg)
	}
	return m, cmds
}
