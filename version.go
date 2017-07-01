package gethweb

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type version struct {}

func (v version) Api() (string) {
	return "0.0.1"
}

func (v version) Node() (string) {
	var result string
	fmt.Println(hp.url)
	resp, err := http.Get(hp.url)
	if err != nil {
		result = "get error"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		result = "get error"
	}
	if result == "" {
		result = string(body)
	}
	return result
}