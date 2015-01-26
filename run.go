package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

var cmdRun = &Command{
	Run:       runRun,
	UsageLine: "run",
	Short:     "runs the configuration",
	Long: ``,
}

func runRun(cmd *Command, args []string) {
	config, err := parseConfiguration()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to retrieve configuration: " + err.Error());
		return
	}

	var finalErr error
	for i := range config.Run {
		cmd := exec.Command("bash", "-c", config.Run[i])

		cmd.Stdin  = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Printf("Running declaration %d: %s\n", i, config.Run[i])
		if err = cmd.Start(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}

		if err = cmd.Wait(); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			finalErr = errors.New("Some command returned an error.")
		}
	}

	if finalErr != nil {
		os.Exit(5)
	}
}
