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
	Long: ``,
}

func runRun(cmd *Command, args []string) {
	config, err := parseConfiguration()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to retrieve configuration: " + err.Error());
	}

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
		}
	}
	_ = config
}
