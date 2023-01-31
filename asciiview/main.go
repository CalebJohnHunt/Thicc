package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	"golang.org/x/term"
)

type viewer struct {
	x, y          int
	width, height int
	image         [][]string
}

func (v *viewer) Init() tea.Cmd {
	return nil
}

func (v *viewer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if msg.Width <= len(v.image[0]) {
			v.width = msg.Width
		}
		if msg.Height < len(v.image) {
			v.height = msg.Height - 1
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return v, tea.Quit
		case "up":
			if v.y > 0 {
				v.y--
			}
		case "down":
			if v.y+v.height < len(v.image) {
				v.y++
			}
		case "left":
			if v.x > 0 {
				v.x--
			}
		case "right":
			if v.x+v.width < len(v.image[0]) {
				v.x++
			}
		}
	}
	return v, nil
}

var styles [5]lipgloss.Style = [...]lipgloss.Style{
	lipgloss.NewStyle().Background(lipgloss.Color("#ff0000")),
	lipgloss.NewStyle().Background(lipgloss.Color("#ff5500")),
	lipgloss.NewStyle().Background(lipgloss.Color("#ffff00")),
	lipgloss.NewStyle().Background(lipgloss.Color("#00ff00")),
	lipgloss.NewStyle().Background(lipgloss.Color("#0000ff"))}

func getStyle(r rune) lipgloss.Style {
	return styles[r%5]
}

func (v *viewer) View() string {
	out := strings.Builder{}
	for _, row := range v.image[v.y : v.y+v.height] {
		for _, letter := range row[v.x : v.x+v.width+1] {
			out.WriteString(letter)
		}
		out.WriteByte('\n')
	}
	out.WriteString(fmt.Sprintf("x: %d y: %d width: %d height: %d imwidth: %d imheight: %d", v.x, v.y, v.width, v.height, len(v.image[0]), len(v.image)))
	return out.String()
}

const asciiImage string = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer tortor ex, lacindo ac leo. Nulla eget orci
ia et euismod id, commodo ac leo. Nulla eget orci molestie ante dapibus convallis. Vivamus euismod purus eni
m, non viverra neque accumsan et. Pellentesque sagittis auctor fringilla. In vitae sodales massa. Cras eu po
suere libero, eget porttitor massa. Donec congue lectus vel rutrum aliquam. Vivamus erat felis, mattis a ull
amcorper a, ornare a sem. Proin molestie pharetra justo, in sagittis lacus pellentesque vitae. Sed molestie 
Phasellus et mi tempus, blandit ligula ac, lobortis leo. Vivamus auctor pretium ornare. Morbi quis mi at ris
us facilisis tincidunt vel sed ante. Nunc tristique ut leo porta faucibus. Cras tincidunt nec nulla sit amet
lementum. Praesent tellus lectus, dapibus quis porta eget, auctor semper nisl. Nullam placerat metus erat,  
at convallis dolor blandit non. Praesent elementum molestie nunc vitae fringilla. Donec bibendum auctor ve  
Maecenas elit leo, cursus id velit vitae, sagittis tempus dui. Aliquam justo turpis, egestas ut aliquam in, 
fringilla sit amet metus. Fusce est enim, feugiat sed tortor in, imperdiet commodo ex. Nulla ornare urna qui
s mollis bibendum. Ut nec nisi malesuada, finibus neque vel, viverra ante. Praesent bibendum quam ac mauris 
fermentum molestie. Pellentesque feugiat mauris sem, vitae tempor diam finibus vitae. Donec at rutrum metus,
 nec pharetra nisi. Aliquam erat volutpat. Curabitur tempus, elit ut volutpat cursus, nulla purus viverra en
Morbi iaculis velit nec leo commodo pharetra. Fusce id vehicula nibh, vitae luctus nisl. Aenean vitae sceler
isque sapien, auctor ullamcorper quam. Integer vitae turpis mattis, dapibus diam quis, tincidunt dui. Pellen
tesque mauris odio, fermentum sed gravida non, aliquet vitae elit. Integer vel libero sapien. Quisque suscip
it lobortis pharetra. Maecenas sed pretium ante. Nullam ac nunc turpis. Curabitur bibendum augue vel hendrer
Mauris pellentesque nunc a turpis cursus feugiat. Donec egestas sit amet elit non consectetur. Quisque nec o
dio ut dolor luctus cursus. Aliquam lobortis urna ac urna congue, pretium bibendum mi congue. Proin viverra 
diam eu scelerisque varius. Proin id auctor turpis. Cras pulvinar rhoncus urna a luctus. Integer eleifend, r
isus et fringilla posuere, nunc ipsum porttitor lectus, at sagittis ante tellus sit amet neque. Pellentesque
Ut varius neque urna, non volutpat augue scelerisque finibus. Aliquam tristique ornare suscipit. Vestibulum 
non tempor nibh. Sed eu dui tempor, posuere erat id, tincidunt odio. Fusce sed sapien commodo, congue magna 
ac, iaculis ex. Integer imperdiet varius augue eget ullamcorper. Aliquam sed pharetra erat. Curabitur commod
o nisl eget massa congue, eget porttitor nisl venenatis. Vestibulum ac placerat libero, nec ornare ante. Dui
Duis sed condimentum est, ac consectetur felis. Cras ut nibh eu odio tristique pulvinar. Integer vestibulum 
odio non est faucibus, eget tincidunt purus fermentum. Fusce lorem felis, aliquet quis nunc quis, pharetra t
empus lorem. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. M
auris dictum egestas aliquam. Cras sed eros nisi. Proin et tortor dolor. Vestibulum molestie tempor purus a 
Mauris suscipit orci dolor, et aliquet diam mattis sed. In sollicitudin dictum neque, non lacinia mi auctor 
in. Integer justo orci, sodales vitae vehicula ac, egestas nec nisl. Phasellus est arcu, placerat a velit vi
tae, viverra tincidunt purus. Curabitur sit amet vehicula eros, id finibus enim. Curabitur egestas mollis ur
na vitae ultrices. Duis hendrerit porttitor erat in viverra. Maecenas ornare aliquet est, viverra sodales to
Duis blandit consectetur enim at fermentum. Curabitur dui orci, lobortis id rhoncus eget, varius id magna. M
auris lacinia sed augue et rutrum. Etiam quis velit risus. Mauris quis justo dui. In quis consectetur lorem.
 Sed dapibus, turpis eu mattis rutrum, mi diam semper turpis, non lacinia nibh felis sit amet ante. Etiam te
 mpus justo eu quam gravida laoreet. Duis volutpat elit leo, in pharetra dui dictum ac. Maecenas urna odio, 
Nullam placerat ultricies lacus id sagittis. Suspendisse potenti. Praesent purus purus, pretium a risus vita
e, euismod tempus mauris. Suspendisse malesuada aliquam diam, eu tempor neque iaculis eget. Proin consectetu
r fermentum sapien accumsan volutpat. Cras facilisis massa ut molestie blandit. Aenean tempor, turpis in lac
inia semper, ipsum neque feugiat massa, ac eleifend est justo non arcu. Maecenas elementum quam at diam semp
ia et euismod id, commodo ac leo. Nulla eget orci molestie ante dapibus convallis. Vivamus euismod purus eni
m, non viverra neque accumsan et. Pellentesque sagittis auctor fringilla. In vitae sodales massa. Cras eu po
suere libero, eget porttitor massa. Donec congue lectus vel rutrum aliquam. Vivamus erat felis, mattis a ull
amcorper a, ornare a sem. Proin molestie pharetra justo, in sagittis lacus pellentesque vitae. Sed molestie 
Phasellus et mi tempus, blandit ligula ac, lobortis leo. Vivamus auctor pretium ornare. Morbi quis mi at ris
us facilisis tincidunt vel sed ante. Nunc tristique ut leo porta faucibus. Cras tincidunt nec nulla sit amet
lementum. Praesent tellus lectus, dapibus quis porta eget, auctor semper nisl. Nullam placerat metus erat,  
at convallis dolor blandit non. Praesent elementum molestie nunc vitae fringilla. Donec bibendum auctor ve  
Maecenas elit leo, cursus id velit vitae, sagittis tempus dui. Aliquam justo turpis, egestas ut aliquam in, 
fringilla sit amet metus. Fusce est enim, feugiat sed tortor in, imperdiet commodo ex. Nulla ornare urna qui
s mollis bibendum. Ut nec nisi malesuada, finibus neque vel, viverra ante. Praesent bibendum quam ac mauris 
fermentum molestie. Pellentesque feugiat mauris sem, vitae tempor diam finibus vitae. Donec at rutrum metus,
 nec pharetra nisi. Aliquam erat volutpat. Curabitur tempus, elit ut volutpat cursus, nulla purus viverra en
Morbi iaculis velit nec leo commodo pharetra. Fusce id vehicula nibh, vitae luctus nisl. Aenean vitae sceler
isque sapien, auctor ullamcorper quam. Integer vitae turpis mattis, dapibus diam quis, tincidunt dui. Pellen
tesque mauris odio, fermentum sed gravida non, aliquet vitae elit. Integer vel libero sapien. Quisque suscip
it lobortis pharetra. Maecenas sed pretium ante. Nullam ac nunc turpis. Curabitur bibendum augue vel hendrer
Mauris pellentesque nunc a turpis cursus feugiat. Donec egestas sit amet elit non consectetur. Quisque nec o
dio ut dolor luctus cursus. Aliquam lobortis urna ac urna congue, pretium bibendum mi congue. Proin viverra 
diam eu scelerisque varius. Proin id auctor turpis. Cras pulvinar rhoncus urna a luctus. Integer eleifend, r
isus et fringilla posuere, nunc ipsum porttitor lectus, at sagittis ante tellus sit amet neque. Pellentesque
Ut varius neque urna, non volutpat augue scelerisque finibus. Aliquam tristique ornare suscipit. Vestibulum 
non tempor nibh. Sed eu dui tempor, posuere erat id, tincidunt odio. Fusce sed sapien commodo, congue magna 
ac, iaculis ex. Integer imperdiet varius augue eget ullamcorper. Aliquam sed pharetra erat. Curabitur commod
o nisl eget massa congue, eget porttitor nisl venenatis. Vestibulum ac placerat libero, nec ornare ante. Dui
Duis sed condimentum est, ac consectetur felis. Cras ut nibh eu odio tristique pulvinar. Integer vestibulum 
odio non est faucibus, eget tincidunt purus fermentum. Fusce lorem felis, aliquet quis nunc quis, pharetra t
empus lorem. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. M
auris dictum egestas aliquam. Cras sed eros nisi. Proin et tortor dolor. Vestibulum molestie tempor purus a 
Mauris suscipit orci dolor, et aliquet diam mattis sed. In sollicitudin dictum neque, non lacinia mi auctor 
in. Integer justo orci, sodales vitae vehicula ac, egestas nec nisl. Phasellus est arcu, placerat a velit vi
tae, viverra tincidunt purus. Curabitur sit amet vehicula eros, id finibus enim. Curabitur egestas mollis ur
na vitae ultrices. Duis hendrerit porttitor erat in viverra. Maecenas ornare aliquet est, viverra sodales to
Duis blandit consectetur enim at fermentum. Curabitur dui orci, lobortis id rhoncus eget, varius id magna. M
auris lacinia sed augue et rutrum. Etiam quis velit risus. Mauris quis justo dui. In quis consectetur lorem.
 Sed dapibus, turpis eu mattis rutrum, mi diam semper turpis, non lacinia nibh felis sit amet ante. Etiam te
 mpus justo eu quam gravida laoreet. Duis volutpat elit leo, in pharetra dui dictum ac. Maecenas urna odio, 
Nullam placerat ultricies lacus id sagittis. Suspendisse potenti. Praesent purus purus, pretium a risus vita
e, euismod tempus mauris. Suspendisse malesuada aliquam diam, eu tempor neque iaculis eget. Proin consectetu
r fermentum sapien accumsan volutpat. Cras facilisis massa ut molestie blandit. Aenean tempor, turpis in lac
inia semper, ipsum neque feugiat massa, ac eleifend est justo non arcu. Maecenas elementum quam at diam semp
ia et euismod id, commodo ac leo. Nulla eget orci molestie ante dapibus convallis. Vivamus euismod purus eni
m, non viverra neque accumsan et. Pellentesque sagittis auctor fringilla. In vitae sodales massa. Cras eu po
suere libero, eget porttitor massa. Donec congue lectus vel rutrum aliquam. Vivamus erat felis, mattis a ull
amcorper a, ornare a sem. Proin molestie pharetra justo, in sagittis lacus pellentesque vitae. Sed molestie 
Phasellus et mi tempus, blandit ligula ac, lobortis leo. Vivamus auctor pretium ornare. Morbi quis mi at ris
us facilisis tincidunt vel sed ante. Nunc tristique ut leo porta faucibus. Cras tincidunt nec nulla sit amet
lementum. Praesent tellus lectus, dapibus quis porta eget, auctor semper nisl. Nullam placerat metus erat,  
at convallis dolor blandit non. Praesent elementum molestie nunc vitae fringilla. Donec bibendum auctor ve  
Maecenas elit leo, cursus id velit vitae, sagittis tempus dui. Aliquam justo turpis, egestas ut aliquam in, 
fringilla sit amet metus. Fusce est enim, feugiat sed tortor in, imperdiet commodo ex. Nulla ornare urna qui
s mollis bibendum. Ut nec nisi malesuada, finibus neque vel, viverra ante. Praesent bibendum quam ac mauris 
fermentum molestie. Pellentesque feugiat mauris sem, vitae tempor diam finibus vitae. Donec at rutrum metus,
 nec pharetra nisi. Aliquam erat volutpat. Curabitur tempus, elit ut volutpat cursus, nulla purus viverra en
Morbi iaculis velit nec leo commodo pharetra. Fusce id vehicula nibh, vitae luctus nisl. Aenean vitae sceler
isque sapien, auctor ullamcorper quam. Integer vitae turpis mattis, dapibus diam quis, tincidunt dui. Pellen
tesque mauris odio, fermentum sed gravida non, aliquet vitae elit. Integer vel libero sapien. Quisque suscip
it lobortis pharetra. Maecenas sed pretium ante. Nullam ac nunc turpis. Curabitur bibendum augue vel hendrer
Mauris pellentesque nunc a turpis cursus feugiat. Donec egestas sit amet elit non consectetur. Quisque nec o
dio ut dolor luctus cursus. Aliquam lobortis urna ac urna congue, pretium bibendum mi congue. Proin viverra 
diam eu scelerisque varius. Proin id auctor turpis. Cras pulvinar rhoncus urna a luctus. Integer eleifend, r
isus et fringilla posuere, nunc ipsum porttitor lectus, at sagittis ante tellus sit amet neque. Pellentesque
Ut varius neque urna, non volutpat augue scelerisque finibus. Aliquam tristique ornare suscipit. Vestibulum 
non tempor nibh. Sed eu dui tempor, posuere erat id, tincidunt odio. Fusce sed sapien commodo, congue magna 
ac, iaculis ex. Integer imperdiet varius augue eget ullamcorper. Aliquam sed pharetra erat. Curabitur commod
o nisl eget massa congue, eget porttitor nisl venenatis. Vestibulum ac placerat libero, nec ornare ante. Dui
Duis sed condimentum est, ac consectetur felis. Cras ut nibh eu odio tristique pulvinar. Integer vestibulum 
odio non est faucibus, eget tincidunt purus fermentum. Fusce lorem felis, aliquet quis nunc quis, pharetra t
empus lorem. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. M
auris dictum egestas aliquam. Cras sed eros nisi. Proin et tortor dolor. Vestibulum molestie tempor purus a 
Mauris suscipit orci dolor, et aliquet diam mattis sed. In sollicitudin dictum neque, non lacinia mi auctor 
in. Integer justo orci, sodales vitae vehicula ac, egestas nec nisl. Phasellus est arcu, placerat a velit vi
tae, viverra tincidunt purus. Curabitur sit amet vehicula eros, id finibus enim. Curabitur egestas mollis ur
na vitae ultrices. Duis hendrerit porttitor erat in viverra. Maecenas ornare aliquet est, viverra sodales to
Duis blandit consectetur enim at fermentum. Curabitur dui orci, lobortis id rhoncus eget, varius id magna. M
auris lacinia sed augue et rutrum. Etiam quis velit risus. Mauris quis justo dui. In quis consectetur lorem.
 Sed dapibus, turpis eu mattis rutrum, mi diam semper turpis, non lacinia nibh felis sit amet ante. Etiam te
 mpus justo eu quam gravida laoreet. Duis volutpat elit leo, in pharetra dui dictum ac. Maecenas urna odio, 
Nullam placerat ultricies lacus id sagittis. Suspendisse potenti. Praesent purus purus, pretium a risus vita
e, euismod tempus mauris. Suspendisse malesuada aliquam diam, eu tempor neque iaculis eget. Proin consectetu
r fermentum sapien accumsan volutpat. Cras facilisis massa ut molestie blandit. Aenean tempor, turpis in lac
inia semper, ipsum neque feugiat massa, ac eleifend est justo non arcu. Maecenas elementum quam at diam semp
ia et euismod id, commodo ac leo. Nulla eget orci molestie ante dapibus convallis. Vivamus euismod purus eni
m, non viverra neque accumsan et. Pellentesque sagittis auctor fringilla. In vitae sodales massa. Cras eu po
suere libero, eget porttitor massa. Donec congue lectus vel rutrum aliquam. Vivamus erat felis, mattis a ull
amcorper a, ornare a sem. Proin molestie pharetra justo, in sagittis lacus pellentesque vitae. Sed molestie 
Phasellus et mi tempus, blandit ligula ac, lobortis leo. Vivamus auctor pretium ornare. Morbi quis mi at ris
us facilisis tincidunt vel sed ante. Nunc tristique ut leo porta faucibus. Cras tincidunt nec nulla sit amet
lementum. Praesent tellus lectus, dapibus quis porta eget, auctor semper nisl. Nullam placerat metus erat,  
at convallis dolor blandit non. Praesent elementum molestie nunc vitae fringilla. Donec bibendum auctor ve  
Maecenas elit leo, cursus id velit vitae, sagittis tempus dui. Aliquam justo turpis, egestas ut aliquam in, 
fringilla sit amet metus. Fusce est enim, feugiat sed tortor in, imperdiet commodo ex. Nulla ornare urna qui
s mollis bibendum. Ut nec nisi malesuada, finibus neque vel, viverra ante. Praesent bibendum quam ac mauris 
fermentum molestie. Pellentesque feugiat mauris sem, vitae tempor diam finibus vitae. Donec at rutrum metus,
 nec pharetra nisi. Aliquam erat volutpat. Curabitur tempus, elit ut volutpat cursus, nulla purus viverra en
Morbi iaculis velit nec leo commodo pharetra. Fusce id vehicula nibh, vitae luctus nisl. Aenean vitae sceler
isque sapien, auctor ullamcorper quam. Integer vitae turpis mattis, dapibus diam quis, tincidunt dui. Pellen
tesque mauris odio, fermentum sed gravida non, aliquet vitae elit. Integer vel libero sapien. Quisque suscip
it lobortis pharetra. Maecenas sed pretium ante. Nullam ac nunc turpis. Curabitur bibendum augue vel hendrer
Mauris pellentesque nunc a turpis cursus feugiat. Donec egestas sit amet elit non consectetur. Quisque nec o
dio ut dolor luctus cursus. Aliquam lobortis urna ac urna congue, pretium bibendum mi congue. Proin viverra 
diam eu scelerisque varius. Proin id auctor turpis. Cras pulvinar rhoncus urna a luctus. Integer eleifend, r
isus et fringilla posuere, nunc ipsum porttitor lectus, at sagittis ante tellus sit amet neque. Pellentesque
Ut varius neque urna, non volutpat augue scelerisque finibus. Aliquam tristique ornare suscipit. Vestibulum 
non tempor nibh. Sed eu dui tempor, posuere erat id, tincidunt odio. Fusce sed sapien commodo, congue magna 
ac, iaculis ex. Integer imperdiet varius augue eget ullamcorper. Aliquam sed pharetra erat. Curabitur commod
o nisl eget massa congue, eget porttitor nisl venenatis. Vestibulum ac placerat libero, nec ornare ante. Dui
Duis sed condimentum est, ac consectetur felis. Cras ut nibh eu odio tristique pulvinar. Integer vestibulum 
odio non est faucibus, eget tincidunt purus fermentum. Fusce lorem felis, aliquet quis nunc quis, pharetra t
empus lorem. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. M
auris dictum egestas aliquam. Cras sed eros nisi. Proin et tortor dolor. Vestibulum molestie tempor purus a 
Mauris suscipit orci dolor, et aliquet diam mattis sed. In sollicitudin dictum neque, non lacinia mi auctor 
in. Integer justo orci, sodales vitae vehicula ac, egestas nec nisl. Phasellus est arcu, placerat a velit vi
tae, viverra tincidunt purus. Curabitur sit amet vehicula eros, id finibus enim. Curabitur egestas mollis ur
na vitae ultrices. Duis hendrerit porttitor erat in viverra. Maecenas ornare aliquet est, viverra sodales to
Duis blandit consectetur enim at fermentum. Curabitur dui orci, lobortis id rhoncus eget, varius id magna. M
auris lacinia sed augue et rutrum. Etiam quis velit risus. Mauris quis justo dui. In quis consectetur lorem.
 Sed dapibus, turpis eu mattis rutrum, mi diam semper turpis, non lacinia nibh felis sit amet ante. Etiam te
 mpus justo eu quam gravida laoreet. Duis volutpat elit leo, in pharetra dui dictum ac. Maecenas urna odio, 
Nullam placerat ultricies lacus id sagittis. Suspendisse potenti. Praesent purus purus, pretium a risus vita
e, euismod tempus mauris. Suspendisse malesuada aliquam diam, eu tempor neque iaculis eget. Proin consectetu
r fermentum sapien accumsan volutpat. Cras facilisis massa ut molestie blandit. Aenean tempor, turpis in lac
inia semper, ipsum neque feugiat massa, ac eleifend est justo non arcu. Maecenas elementum quam at diam semp
`

func makeimage() [][]string {
	image := make([][]string, 0)
	row := make([]string, 0)
	for _, r := range asciiImage {
		if r == '\n' {
			image = append(image, row)
			row = make([]string, 0)
			continue
		}
		row = append(row, getStyle(r).Render(" "))
	}
	return image
}

func main() {
	image := makeimage()

	v := &viewer{image: image}

	width, height, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Couldn't get width and height. Using default (5).")
		v.width = 5
		v.height = 5 - 1
	} else {
		fmt.Printf("Width: %d Height: %d (press enter to continue)", width, height)
		if w := len(v.image[0]); width > w {
			v.width = w
		} else {
			v.width = width
		}
		if h := len(v.image); height > h {
			v.height = h
		} else {
			v.height = height
		}
		v.height--
	}

	fmt.Scanln()

	// tea.NewProgram(v, tea.WithAltScreen()).Run()

	s, err := wish.NewServer(wish.WithMiddleware(bm.Middleware(teaHandler)), wish.WithAddress(":6969"))
	if err != nil {
		panic(err)
	}
	err = s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, active := s.Pty()
	if !active {
		wish.Fatalln(s, "no active terminal, skipping")
		return nil, nil
	}
	m := viewer{
		image:  makeimage(),
		width:  pty.Window.Width,
		height: pty.Window.Height,
	}
	return &m, []tea.ProgramOption{tea.WithAltScreen()}
}
