package e2e

import (
	"os/exec"
	"testing"
)

// If no rules are provided watchcat only checks if a directory exists or not.
func TestWatchcatCliReportModeNoRules(t *testing.T) {
	setup(t)
	cmd := exec.Command("./../build/watchcat", "-config", "../e2e/cli01/cfg", "-mode", "report-rules")
	expected := // expected program output.
	" [ ] main (config)          |" + NL +
		"   [X] firefox              |" + NL +
		"   [X] downloads            |" + NL +
		"   [ ] downloads/notthere   |x" + NL +
		"   [X] downloads/notthere2  |" + NL +
		"   [X] downloads/done       |" + NL +
		"   [X] downloads/incomplete |" + NL +
		" [X] secondary (Watchconf)  |" + NL +
		"   [X] firefox              |e" + NL +
		"   [X] downloads            |" + NL +
		"   [X] downloads/done       |eF" + NL +
		"   [X] downloads/incomplete |e"
	actual := captureAndExecute(t, cmd)
	verify(t, actual, expected)
	cleanup(t)
}
