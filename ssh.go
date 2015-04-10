package main

import (
	"fmt"
	"os"
	"os/exec"
)

var cmdSsh = &Command{
	Run:       runSsh,
	UsageLine: "ssh",
	Short:     "connects to Coduno via ssh",
	Long:      `This is helpful if you want to check whether Coduno has your public key.`,
}

func runSsh(cmd *Command, args []string) {
	c := exec.Command("ssh", "git@git.cod.uno")

	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	if err := c.Wait(); err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
}
