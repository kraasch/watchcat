package main

import (
	"flag"
	"fmt"
	"os"

	// for making a nice centred box.
	tea "github.com/charmbracelet/bubbletea"
	lip "github.com/charmbracelet/lipgloss"

	// local packages.
	wcat "github.com/kraasch/watchcat/pkg/wcat"
)

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
		str = wcat.Toast("Hello!")
	} else {
		str = wcat.Toast("Hi!")
	}
	str = styleBox.Render(str)
	return lip.Place(m.width, m.height, lip.Center, lip.Center, str)
}

func main() {
	// parse flags.
	flag.BoolVar(&verbose, "verbose", false, "Show info")
	flag.BoolVar(&suppress, "suppress", false, "Print nothing")
	configFlag := flag.String("config", "", "Path to the configuration file.") // TODO: insert good default value.
	// Possible modes are: tui, list, print-config, report-rules.
	modeFlag := flag.String("mode", "none", "Mode of operation (e.g. tui, list, print-config)") // TODO: insert good default value.
	flag.Parse()
	configStr := *configFlag
	modeStr := *modeFlag
	// create watchcat.
	wc := wcat.New()
	wc.ReadConfig(configStr)
	// evaluate mode and what to do.
	switch modeStr { // if mode != "tui" then solve things on the cli.
	case "tui": // launch tui.
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
	case "list": // solve on cli.
		fmt.Printf("%s", wc.ListTargets())
	case "print-config": // solve on cli.
		fmt.Printf("%s", wc.PrintConfig())
	case "report-rules": // solve on cli.
		fmt.Printf("%s", wc.PrintConfig())
	}
} // fin.
