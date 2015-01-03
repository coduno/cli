package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

type Challenge struct {
	Archived             bool        `json:"archived"`
	CreatedAt            string      `json:"created_at"`
	DefaultBranch        string      `json:"default_branch"`
	Description          interface{} `json:"description"`
	HttpUrlToRepo        string      `json:"http_url_to_repo"`
	ID                   float64     `json:"id"`
	IssuesEnabled        bool        `json:"issues_enabled"`
	LastActivityAt       string      `json:"last_activity_at"`
	MergeRequestsEnabled bool        `json:"merge_requests_enabled"`
	Name                 string      `json:"name"`
	NameWithNamespace    string      `json:"name_with_namespace"`
	Namespace            struct {
		CreatedAt   string  `json:"created_at"`
		Description string  `json:"description"`
		ID          float64 `json:"id"`
		Name        string  `json:"name"`
		OwnerID     float64 `json:"owner_id"`
		Path        string  `json:"path"`
		UpdatedAt   string  `json:"updated_at"`
	} `json:"namespace"`
	Owner struct {
		CreatedAt string  `json:"created_at"`
		ID        float64 `json:"id"`
		Name      string  `json:"name"`
	} `json:"owner"`
	Path              string  `json:"path"`
	PathWithNamespace string  `json:"path_with_namespace"`
	Public            bool    `json:"public"`
	SnippetsEnabled   bool    `json:"snippets_enabled"`
	SshUrlToRepo      string  `json:"ssh_url_to_repo"`
	VisibilityLevel   float64 `json:"visibility_level"`
	WebURL            string  `json:"web_url"`
	WikiEnabled       bool    `json:"wiki_enabled"`
}

var cmdChallenge = &Command{
	Run:       getChallenge,
	UsageLine: "challenge <challenge-name>",
	Short:     "fetch challenge from server",
	Long: `
Fetch challenge <challenge-name> from server.

All needed files will be saved in a directory named <challende-name>.
	`,
}

func getChallenge(cmd *Command, args []string) {
	if len(args) < 1 {
		// No argument given, list all available challenges
		listChallenges(cmd, args)
	} else if len(args) == 1 {
		// challenge name given, clone into directory
		cloneChallenge(cmd, args)
	} else {
		// Invalid number of arguments, show usage
		fmt.Fprintln(os.Stderr, "Invalid number of arguments")
		os.Exit(2)
	}
}

func listChallenges(cmd *Command, args []string) {
	auth := url.Values{}
	auth.Add("login", "dummy")
	auth.Add("password", "dummydummy")
	req, err := http.NewRequest("POST", "http://git.cod.uno/api/v3/session", strings.NewReader(auth.Encode()))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusCreated { // All good
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		var dat map[string]interface{}
		if err := json.Unmarshal(body, &dat); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		token := dat["private_token"].(string)
		req, err = http.NewRequest("GET", "http://git.cod.uno/api/v3/projects", nil)
		req.Header = map[string][]string{
			"PRIVATE-TOKEN": {token},
		}
		res, err = client.Do(req)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		defer res.Body.Close()
		if res.StatusCode == http.StatusOK { // Read list
			body, err = ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}

			challenges := make([]Challenge, 0)
			if err = json.Unmarshal(body, &challenges); err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			for i := 0; i < len(challenges); i++ {
				fmt.Printf("%s: %s\n", challenges[i].Path, challenges[i].Description)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Something went wrong while fetching challenge list (%d)\n", res.StatusCode)
			os.Exit(1)
		}
	} else {
		fmt.Fprintln(os.Stderr, "Something went wrong while authorizing access")
		os.Exit(1)
	}
}

func cloneChallenge(cmd *Command, args []string) {
	repo := "http://git.cod.uno/challenges/" + args[0] + ".git"
	c := exec.Command("git", "clone", repo)
	outstr, err := c.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to fetch challenge <"+args[0]+">.")
		fmt.Fprintln(os.Stderr, "Reason: "+string(outstr)) // TODO: Remove message from production tool
		os.Exit(2)
	} else {
		fmt.Println(os.Stdout, "Successfully fetched into directory "+args[0])
	}
}
