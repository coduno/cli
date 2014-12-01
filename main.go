package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var url = "http://coduno.appspot.com/api/hello"

func main() {
	fmt.Println("hello world")
	hello()
}

func hello() {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
