package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/coduno/netrc"
	"github.com/howeyc/gopass"
	"github.com/mitchellh/go-homedir"
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
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password := string(gopass.GetPasswd())
	password = strings.TrimSpace(password)

	// Get pub key TODO: allow for use of custom key location
	path, _ := homedir.Expand("~/.ssh/id_rsa.pub")
	keyfile, err := os.Open(path)
	if err != nil {
		fmt.Printf("Please provide your public RSA key at %s\n", path)
		os.Exit(1)
	}

	authorization := "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))
	req, err := http.NewRequest("POST", "https://coduno.appspot.com/api/token", keyfile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	req.Header = map[string][]string{
		"Authorization": {authorization},
		"Connection":    {"close"},
	}

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode == http.StatusOK { // All good
		netrcx, err := netrc.Parse()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read netrc: %s", err.Error())
		}

		netrcx.Entries["git.cod.uno"] = netrc.Entry{Login: username, Password: string(body)}

		err = netrcx.Save()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to save netrc: %s", err.Error())
		}
	} else {
		fmt.Print(string(body))
	}
}
