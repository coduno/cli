package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/oauth2"

	"github.com/coduno/captain"
	"github.com/coduno/netrc"
	"github.com/howeyc/gopass"
)

var cmdAuth = &captain.Command{
	Run: func(cmd []*captain.Command, args []string) {
		fmt.Println("ibc")
	},
	UsageLine: "auth",
	Short:     "authentication",
	Long:      "Let's you log in and out of Coduno.",
	SubCommands: []*captain.Command{
		&captain.Command{
			Run:       runAuthLogin,
			UsageLine: "login",
			Short:     "Authenticate with Coduno",
			Long: `
Guides you through an OAuth flow to obtain an authentication token from Coduno.

The authentication token will be saved to ` + netrc.Location() + `.
			`,
		},
		&captain.Command{
			Run:       runAuthLogout,
			UsageLine: "logout",
			Short:     "Delete all credentials",
			Long:      "Deletes your current authorization token from " + netrc.Location() + ".",
		},
		&captain.Command{
			Run:       runAuthStatus,
			UsageLine: "status",
			Short:     "Tells you whether you are logged in",
			Long:      "Looks for credentials in " + netrc.Location() + " and checks whether you are logged in.",
			Flag:      flagAuthStatus,
		},
	},
}

var flagAuthStatus = flag.NewFlagSet("", flag.PanicOnError)
var authStatusOffline = false

func init() {
	flagAuthStatus.BoolVar(&authStatusOffline, "offline", false, "only check for local credentials and assume they are valid")
}

func runAuthLogout(cmd []*captain.Command, args []string) {
	entries, err := netrc.Parse()
	if err != nil {
		fatalf(1, "status: %s", err.Error())
	}
	delete(entries, "api.cod.uno")
	entries.Save()
}

func runAuthStatus(cmd []*captain.Command, args []string) {
	entries, err := netrc.Parse()
	if err != nil {
		fatalf(1, "status: %s", err.Error())
	}
	entry, ok := entries["api.cod.uno"]
	if !ok {
		fmt.Println("You are not authenticated.")
		return
	}
	if authStatusOffline {
		fmt.Println("You are authenticated!")
		return
	}
	panic("not implemented")
	fmt.Println("You are not authenticated.")
	_ = entry
}

func runAuthLogin(cmd []*captain.Command, args []string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password := string(gopass.GetPasswd())
	password = strings.TrimSpace(password)

	c := oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Endpoint: oauth2.Endpoint{
			AuthURL:  build("/oauth/auth"),
			TokenURL: build("/oauth/token"),
		},
		Scopes: nil,
	}

	_, err := c.PasswordCredentialsToken(nil, username, password)
	if err != nil {
		panic(err)
	}

	// Get pub key TODO: allow for use of custom key location
	// path, _ := homedir.Expand("~/.ssh/id_rsa.pub")
	// keyfile, err := os.Open(path)
	// if err != nil {
	// 	fmt.Printf("Please provide your public RSA key at %s\n", path)
	// 	os.Exit(1)
	// }

	// 	authorization := "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))
	//
	// 	req, err := http.NewRequest("POST", build+"/tokens", keyfile)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		os.Exit(1)
	// 	}
	// 	req.Header = map[string][]string{
	// 		"Authorization": {authorization},
	// 		"Connection":    {"close"},
	// 	}
	//
	// 	client := new(http.Client)
	// 	res, err := client.Do(req)
	// 	if err != nil {
	// 		fmt.Printf(err.Error())
	// 	}
	// 	defer res.Body.Close()
	//
	// 	body, err := ioutil.ReadAll(res.Body)
	//
	// 	if res.StatusCode == http.StatusOK { // All good
	// 		netrcx, err := netrc.Parse()
	//
	// 		if err != nil {
	// 			fmt.Fprintf(os.Stderr, "Failed to read netrc: %s", err.Error())
	// 		}
	//
	// 		netrcx["git.cod.uno"] = netrc.Entry{Login: username, Password: string(body)}
	//
	// 		err = netrcx.Save()
	// 		if err != nil {
	// 			fmt.Fprintf(os.Stderr, "Failed to save netrc: %s", err.Error())
	// 		}
	// 	} else {
	// 		fmt.Print(string(body))
	// 	}
}
