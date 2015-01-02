package main

import (
	"./netrc"
	"fmt"
)

var cmdLogin = &Command{
	Run:       runLogin,
	UsageLine: "login",
	Short:     "log in with Coduno",
	Long: `
Guides you through an OAuth flow to obtain an authentication token from Coduno.

The authentication token will be saved to` + netrc.Location() + `.
	`,
}

func runLogin(cmd *Command, args []string) {
	netrcx, _ := netrc.Parse()
	fmt.Printf("%d", len(netrcx.Entries))
	netrcx.Save()
}
