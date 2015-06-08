package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var stats = false

func init() {
	cmdRun.Flag.BoolVar(&stats, "stats", false,
		"Enable stats tracking to file stats.log")
}

var cmdRun = &Command{
	Run:         runRun,
	UsageLine:   "run",
	Short:       "runs the configuration",
	Long:        ``,
	Flag:        flag.NewFlagSet("run", flag.ExitOnError),
	CustomFlags: false,
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

	if stats {
		usage := c.ProcessState.SysUsage()
		f, err := os.Create("stats.log")
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}

		defer f.Close()
		enc := json.NewEncoder(f)
		enc.Encode(usage)
	}
}
