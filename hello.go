package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var cmdHello = &Command{
	Run:       runHello,
	UsageLine: "hello",
	Short:     "test connection to Coduno API",
	Long: `
hello reaches out to the API and waits for a reply.
	`,
}

var url = "http://coduno.appspot.com/api/hello"

func runHello(cmd *Command, args []string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
