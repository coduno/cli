package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

var cmdPush = &Command{
	Run:       runPush,
	UsageLine: "push",
	Short:     "pushes your results to Coduno",
	Long:      ``,
}

func runPush(cmd *Command, args []string) {
	repoBytes, err := exec.Command("git", "rev-parse", "--show-toplevel").CombinedOutput()

	if err != nil {
		fmt.Fprintln(os.Stderr, "I don't know what to push.")
		os.Exit(2)
	}

	repo := path.Base(strings.Trim(string(repoBytes), "\r\n"))

	remote := "git@git.cod.uno:" + repo
	push := exec.Command("git", "push", remote)

	push.Stdin = os.Stdin
	push.Stdout = os.Stdout
	push.Stderr = os.Stderr

	if err = push.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	if err = push.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
