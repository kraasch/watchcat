package e2e

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

const (
	softHighlight   = "\x1b[1;38;2;100;100;100m" // ANSI foreground color (= gray).
	strongHighlight = "\x1b[1;38;2;255;0;0m"     // ANSI foreground color (= red).
	warningColor    = "\x1b[48;5;56m"            // ANSI background color (= purple).
	resetColor      = "\x1b[0m"                  // ANSI clear formatting.
)

var NL = fmt.Sprintln()

func highlightSpaces(input string) string {
	replaced := input
	// things you usually do not want to see in your strings.
	replaced = strings.ReplaceAll(replaced, "\n", softHighlight+"¶"+resetColor+NL) // line breaks (linux/mac)
	replaced = strings.ReplaceAll(replaced, " ", strongHighlight+"-"+resetColor)   // spaces.
	replaced = strings.ReplaceAll(replaced, "\t", strongHighlight+"++"+resetColor) // tabs.
	// things you usually do not want to see in your strings.
	replaced = strings.ReplaceAll(replaced, "\r\n", warningColor+"¶"+resetColor+NL)    // line breaks (windows).
	replaced = strings.ReplaceAll(replaced, "\r", warningColor+"<ret>"+resetColor)     // carriage return.
	replaced = strings.ReplaceAll(replaced, "\v", warningColor+"<vtab>"+resetColor)    // vertical tabs.
	replaced = strings.ReplaceAll(replaced, "\f", warningColor+"<feed>"+resetColor)    // form feed.
	replaced = strings.ReplaceAll(replaced, "\u00A0", warningColor+"<nbs>"+resetColor) // non-breaking space.
	return "'" + replaced + "'"
}

func captureAndExecute(t *testing.T, cmd *exec.Cmd) string {
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to run command: %v, stderr: %s", err, stderr.String())
	}
	actual := stdout.String()
	return actual
}

func verify(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Errorf("Output mismatch.\nExpected: %s\nGot: %s", expected, actual)
	}
}

func cleanup(t *testing.T) {
	cmdClean := exec.Command("rm", "-r", "./cli01/")
	errClean := cmdClean.Run()
	if errClean != nil {
		t.Fatalf("Failed directory clean up: '%v'", errClean)
	}
}

func setup(t *testing.T) {
	cmdCreate := exec.Command("bash", "./cli01_create.sh")
	errCreate := cmdCreate.Run()
	if errCreate != nil {
		t.Fatalf("Failed directory creation: '%v'", errCreate)
	}
}

func TestWatchcatCliListCommand(t *testing.T) {
	setup(t)
	cmd := exec.Command("./../build/watchcat", "-config", "../e2e/cli01/cfg", "-list")
	expected := // expected program output.
	"main" + NL +
		"secondary" + NL
	actual := captureAndExecute(t, cmd)
	verify(t, highlightSpaces(actual), highlightSpaces(expected))
	cleanup(t)
}

// func TestWatchcatCliPrintConfigCommand(t *testing.T) {
// 	setup(t)
// 	cmd := exec.Command("./../build/watchcat", "-config", "../e2e/cli01/cfg", "-list")
// 	expected := // expected program output.
// 	"main" + NL +
// 		"secondary"
// 	actual := captureAndExecute(t, cmd)
// 	verify(t, actual, expected)
// 	cleanup(t)
// }
