// Package wcat contains the main program under the TUI and under the CLI.
package wcat

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/kraasch/watchcat/pkg/gocfg"
)

const (
	value                    = "Toast: " // TODO: remove later.
	defaultWatchconfFilename = ".Watchconf"
	defaultConfigFilename    = "config.toml"
)

var NL = fmt.Sprintln()

type Watchcat struct {
	config gocfg.WatchcatConfig
	toml   gocfg.SerializedConfig
}

func New() Watchcat {
	return Watchcat{}
}

func Toast(in string) string {
	return fmt.Sprintf("%#v", value+in)
}

func (w *Watchcat) deepRead() error {
	for i := range w.toml.Targets {
		if w.toml.Targets[i].RulesLocation == "wc" {
			err := w.readWatchconf(&w.toml.Targets[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (w *Watchcat) ReadConfig(path string) { // TODO: add error handling.
	file := defaultConfigFilename
	w.config = gocfg.WatchcatConfig{Filename: file, Path: path}
	result, err0 := w.config.ParseToml()
	w.toml = result
	if err0 != nil {
		fmt.Println("Error reading config file:", err0) // TODO: do propper error handling here.
	}
	// Do a deep read and find the Watchconf files for the targets which have no rules in the config file.
	err1 := w.deepRead()
	if err1 != nil {
		fmt.Println("Error reading one of the watchconf files:", err1) // TODO: do propper error handling here.
	}
}

func (w *Watchcat) ListTargets() string {
	var sb strings.Builder
	for _, target := range w.toml.Targets {
		sb.WriteString(target.Name + NL)
	}
	return sb.String()
}

func indentMultiLine(s string, indent string) string {
	lines := strings.Split(s, NL)
	for i, line := range lines {
		lines[i] = indent + line
	}
	return strings.Join(lines, NL)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (w *Watchcat) readWatchconf(target *gocfg.Target) error {
	var path string
	path = gocfg.InsertSeparator(w.config.Path, target.Dir)      // path to config + parsed path to watchconf.
	path = gocfg.InsertSeparator(path, defaultWatchconfFilename) // add watchconf filename.
	// check if file exists.
	if !fileExists(path) {
		return errors.New("Watchconf not found at '" + path + "'")
	}
	// attempt read.
	result, err := gocfg.ReadRawText(path)
	if err != nil {
		return err
	}
	// successful read.
	target.Rules = result
	return nil
}

// padLinesToLongest adds padding for each line.
func padLinesToLongest(input string) string {
	lines := strings.Split(input, NL)
	// Find the line with the maximum length.
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	// Pad each line with spaces to match the maximum.
	for i, line := range lines {
		padding := maxLen - len(line)
		lines[i] = line + strings.Repeat(" ", padding+1) + "|"
	}
	// Join the lines back into a single string.
	return strings.Join(lines, NL)
}

func (w *Watchcat) PrintConfig() string {
	var sb strings.Builder
	length := len(w.toml.Targets)
	for i, target := range w.toml.Targets {
		mode := ""
		switch target.RulesLocation {
		case "cfg":
			mode = "config"
		case "wc":
			mode = "Watchconf"
		}
		rules := target.Rules
		rules = strings.TrimRight(rules, NL)
		sb.WriteString(target.Name + " (" + mode + ")" + NL)
		indentedRules := indentMultiLine(rules, "  ") // Two spaces.
		sb.WriteString(indentedRules)
		if i < length-1 { // Do not add a new line after the last target.
			sb.WriteString(NL)
		}
	}
	return padLinesToLongest(sb.String())
}

func (w *Watchcat) ReportRules() string {
	var sb strings.Builder
	length := len(w.toml.Targets)
	for i, target := range w.toml.Targets {
		mode := ""
		switch target.RulesLocation {
		case "cfg":
			mode = "config"
		case "wc":
			mode = "Watchconf"
		}
		rules := target.Rules
		rules = strings.TrimRight(rules, NL)
		sb.WriteString(target.Name + " (" + mode + ")" + NL)
		indentedRules := indentMultiLine(rules, "  ") // Two spaces.
		sb.WriteString(indentedRules)
		if i < length-1 { // Do not add a new line after the last target.
			sb.WriteString(NL)
		}
	}
	return padLinesToLongest(sb.String())
}
