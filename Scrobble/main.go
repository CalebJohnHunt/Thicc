package main

import (
	"bufio"
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	rs *rand.Rand
)

func main() {
	f, err := os.Create("log")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)

	width := flag.Int("width", 20, "Width of the playing grid.")
	height := flag.Int("height", 20, "Height of the playing grid.")
	spaceString := flag.String("space", ".", "The empty space character. Only the first character of the supplied string will be used.")
	lettersString := flag.String("letters", "", "Use these letters instead of randomly generating them.")
	minVowels := flag.Int("minVowels", 5, "Minimum number of vowels.  If -minVowels >= -totalLetters, all letters will be vowels.")
	totalLetters := flag.Int("totalLetters", 20, "Total number of letters. If -minVowels >= -totalLetters, all letters will be vowels.")
	seed := flag.Int("seed", -1, "Random seed (-1 means random).")
	flag.Parse()
	if *seed == -1 {
		*seed = int(time.Now().UnixNano())
	}
	rs = rand.New(rand.NewSource(int64(*seed)))
	if len(*spaceString) > 0 {
		space = rune((*spaceString)[0])
	}

	letters := map[rune]int{}
	if len(*lettersString) > 0 {
		for _, r := range *lettersString {
			letters[r]++
		}
	} else {
		letters = makeRandomPool(*totalLetters - *minVowels)
		addVowels(*minVowels, letters)
	}

	p := tea.NewProgram(
		&model{
			c:          cursor{sleep: time.Second},
			body:       makeBody(*width, *height),
			colors:     makeColors(*width, *height),
			dictionary: makeDictionary(),
			dir:        RIGHT,
			letters:    letters,
		}, tea.WithAltScreen())
	p.Run()
}

func makeDictionary() (dict map[string]struct{}) {
	dict = map[string]struct{}{}
	f, err := os.Open("dictionary.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		dict[scanner.Text()] = struct{}{}
	}
	return dict
}

func makeColors(width, height int) [][]colorDecider {
	colors := make([][]colorDecider, height)
	for i := range colors {
		colors[i] = make([]colorDecider, width)
	}
	return colors
}

func makeBody(width, height int) [][]rune {
	body := make([][]rune, height)
	for i := range body {
		body[i] = make([]rune, width)
		for j := range body[i] {
			body[i][j] = space
		}
	}
	return body
}

var vowels [5]rune = [...]rune{'A', 'E', 'I', 'O', 'U'}

func addVowels(n int, m map[rune]int) {
	for i := 0; i < n; i++ {
		m[vowels[rs.Int31n(5)]]++
	}
}

func makeRandomPool(n int) map[rune]int {
	ans := map[rune]int{}
	for i := 0; i < n; i++ {
		l := rune(rs.Int31n(26) + 'A')
		if l == 'Q' {
			if i != n-1 {
				ans['U']++
				i++
			} else {
				ans[l]--
				i--
			}
		}

		ans[l]++
	}
	return ans
}
