package main

import (
	"fmt"
	"os"
	"os/exec"
)

var cmdPrepare = &Command{
	Run:       runPrepare,
	UsageLine: "prepare",
	Short:     "prepares everything for building and running",
	Long: ``,
}

func runPrepare(cmd *Command, args []string) {
	config, err := parseConfiguration()

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error());
		return
	}

	for i := range config.Prepare {
		c := exec.Command("bash", "-c", config.Prepare[i])

		c.Stdin  = os.Stdin
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
}
