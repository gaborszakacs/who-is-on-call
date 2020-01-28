// +build e2e

package e2e

import (
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestE2E(t *testing.T) {
	err := os.Chdir("..")
	if err != nil {
		t.Fatalf("Failed to changed dir: %v", err)
	}
	timeStart := time.Now()
	cmd := exec.Command("go", "run", "main.go")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run: %v: %s", err, string(output))
	}

	got := string(output)
	want := `So the one
who is on-call this week
is...
JOZSI`

	if got != want {
		t.Errorf("got:\n%s\n\nwant:\n%s", got, want)
	}

	if time.Since(timeStart) < 5*time.Second {
		t.Errorf("should have lasted at least 5 sec")
	}
}
