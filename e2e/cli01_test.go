package e2e

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestWatchcatCliOutput(t *testing.T) {
	// Prepare the command: 'watchcat --level 3'
	cmd := exec.Command("./../build/watchcat", "-config", "../e2e/cli01/cfg", "-list")
	// Capture stdout and stderr.
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	// Run the command.
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to run command: %v, stderr: %s", err, stderr.String())
	}
	actual := stdout.String()
	expected := "bonkers!"
	// Verify the actual output matches the expected.
	if actual != expected {
		t.Errorf("Output mismatch.\nExpected: %s\nGot: %s", expected, actual)
	}
}
