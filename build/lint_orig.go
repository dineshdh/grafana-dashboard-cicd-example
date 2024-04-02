package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var lint_orig = goyek.Define(goyek.Task{
	Name:  "lint_orig",
	Usage: "Original attempt to lint dashboards",
	Action: func(a *goyek.A) {
		err := filepath.Walk("../dashboards",
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
					cmd.Exec(a, fmt.Sprintf("dashboard-linter lint %s --verbose --strict", path), cmd.Dir("."))
				}
				return nil
			})
		if err != nil {
			fmt.Println("Error walking the directory:", err)
		}
	},
})
