package main

import (
	"./netrc"
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/howeyc/gopass"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
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

	authorization := "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))
	req := new(http.Request)
	req.URL, _ = url.Parse("http://coduno.appspot.com/api/token")
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
			fmt.Printf("Failed to read netrc:", err.Error())
		}

		netrcx.Entries["git.cod.uno"] = netrc.Entry{Login: username, Password: string(body)}

		err = netrcx.Save()
		if err != nil {
			fmt.Printf("Failed to save netrc:", err.Error())
		}
	} else {
		fmt.Print(string(body))
	}
}
