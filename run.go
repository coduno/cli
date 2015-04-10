package main

import (
	"fmt"
	"os"
	"os/exec"
)

var cmdRun = &Command{
	Run:       runRun,
	UsageLine: "run",
	Short:     "runs the configuration",
	Long:      ``,
}

func runRun(cmd *Command, args []string) {
	config, err := parseConfiguration()

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	c := exec.Command("bash", "-c", config.Run)

	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err = c.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	if err = c.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
