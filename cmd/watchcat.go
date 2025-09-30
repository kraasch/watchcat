package main

import (
	"flag"
	"fmt"
	"os"

	// for making a nice centred box.
	tea "github.com/charmbracelet/bubbletea"
	lip "github.com/charmbracelet/lipgloss"

	// local packages.
	gocfg "github.com/kraasch/watchcat/pkg/gocfg"
	engine "github.com/kraasch/watchcat/pkg/wcat"
)

const (
	defaultConfigFilename = "config.toml"
)

var NL = fmt.Sprintln()

var (
	// return value.
	tuiOutput = ""
	// flags.
	verbose  = false
	suppress = false
	// styles.
	styleBox = lip.NewStyle().
			BorderStyle(lip.NormalBorder()).
			BorderForeground(lip.Color("56"))
)

type model struct {
	width  int
	height int
}

func (m model) Init() tea.Cmd {
	return func() tea.Msg { return nil }
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			tuiOutput = "You quit on me!"
			return m, tea.Quit
		}
	}
	return m, cmd
}

func (m model) View() string {
	var str string
	if verbose {
		str = engine.Toast("Hello!")
	} else {
		str = engine.Toast("Hi!")
	}
	str = styleBox.Render(str)
	return lip.Place(m.width, m.height, lip.Center, lip.Center, str)
}

func main() {
	// parse flags.
	var list bool // alias flag for '-mode = "list"'
	flag.BoolVar(&verbose, "verbose", false, "Show info")
	flag.BoolVar(&suppress, "suppress", false, "Print nothing")
	flag.BoolVar(&list, "list", false, "List mode (alias).")
	configFlag := flag.String("config", "", "Path to the configuration file.")       // TODO: insert good default value.
	modeFlag := flag.String("mode", "", "Mode of operation (e.g., none, list, tui)") // TODO: insert good default value.
	flag.Parse()
	configStr := *configFlag
	modeStr := *modeFlag
	if list && modeStr != "" { // do not accept mode and list flag at the same time.
		fmt.Println("Flag error: only use -list or -mode 'mode', not both.")
		os.Exit(1)
	} else if list && modeStr == "" {
		modeStr = "list"
	}

	// pase config file.
	cfg := gocfg.Config{Filename: defaultConfigFilename, Path: configStr}

	switch modeStr {
	case "tui":
		// init model.
		m := model{0, 0}
		// start bubbletea.
		if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		// print the last highlighted value in calendar to stdout.
		if !suppress {
			fmt.Println(tuiOutput)
		}
	case "list": // if mode != "tui" then solve things on the cli.
		txt, err := cfg.ReadRawText()
		if err != nil {
			fmt.Println("Error reading config file:", err)
			os.Exit(1)
		}
		fmt.Printf("config: %s\n", txt)
		fmt.Printf("values: apples\npears\nbananas\ngrapes\n")
	}
} // fin.
