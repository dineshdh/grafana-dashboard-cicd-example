// TestDashboardLinting is a test function to lint Grafana dashboard JSON files.
package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestDashboardLinting(t *testing.T) {
	// Walk through the "../dashboards" directory and its subdirectories
	err := filepath.Walk("../dashboards",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// Check if the file is not a directory and has a ".json" extension
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
				t.Run(path, func(t *testing.T) {
					t.Parallel()                       // Run tests in parallel, so they don't block each other and it's faster
					time.Sleep(100 * time.Millisecond) // Simulate a slow test

					t.Logf("Linting Grafana dashboard JSON file: %s", path)
					// Run the "dashboard-linter" command to lint the JSON file
					cmd := exec.Command("dashboard-linter", "lint", path, "--verbose", "--strict")
					output, err := cmd.CombinedOutput()
					if err != nil {
						t.Errorf("Error running dashboard-linter: %v", err)
					}
					if testing.Verbose() {
						t.Logf("\n%s", output)
					}
				})
			}
			return nil
		})
	if err != nil {
		t.Errorf("Error walking the directory: %v", err)
	}
}
