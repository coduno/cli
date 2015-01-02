package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

var cmdVersion = &Command{
	Run:       runVersion,
	UsageLine: "version",
	Short:     "print Coduno version",
	Long: `
	Version prints all necessary version information needed to
	trace bugs. Include its output when filing a bug.
	`,
}

var Version = ""
var BuildTime = ""

func runVersion(cmd *Command, args []string) {
	fmt.Println("coduno version", Version, BuildTime)
	fmt.Println("go version", runtime.Version())

	out, err := exec.Command("git", "version").Output()
	if err == nil {
		fmt.Println(strings.TrimSpace(string(out)))
	}
}
