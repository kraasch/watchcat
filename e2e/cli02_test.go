package e2e

import (
	"os/exec"
	"testing"
)

func TestWatchcatCliListCommand2(t *testing.T) {
	setup(t)
	cmd := exec.Command("./../build/watchcat", "-config", "../e2e/cli01/cfg", "-mode", "list")
	expected := // expected program output.
	"main" + NL +
		"secondary" + NL
	actual := captureAndExecute(t, cmd)
	verify(t, actual, expected)
	cleanup(t)
}

func TestWatchcatCliPrintConfigCommand2(t *testing.T) {
	setup(t)
	cmd := exec.Command("./../build/watchcat", "-config", "../e2e/cli01/cfg", "-mode", "print-config")
	expected := // expected program output.
	"main (config)          |" + NL +
		"  firefox              |" + NL +
		"  downloads            |" + NL +
		"  downloads/done       |" + NL +
		"  downloads/incomplete |" + NL +
		"secondary (Watchconf)  |" + NL +
		"  firefox              |" + NL +
		"  downloads            |" + NL +
		"  downloads/done       |" + NL +
		"  downloads/incomplete |"
	actual := captureAndExecute(t, cmd)
	verify(t, actual, expected)
	cleanup(t)
}
