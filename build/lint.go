package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var lint = goyek.Define(goyek.Task{
	Name:  "lint",
	Usage: "Preferred way to lint dashboards",
	Action: func(a *goyek.A) {
		cmd.Exec(a, "go test dashboards_test.go -v", cmd.Dir("../dashboards"))
	},
})
