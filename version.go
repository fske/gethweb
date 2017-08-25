package gethweb

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type version struct {
	httpPrvd httpProvider
}

func (v version) Api() (string) {
	return "0.0.1"
}

func (v version) Node() (string) {
	var result string
	fmt.Println(v.httpPrvd.url)
	resp, err := http.Get(v.httpPrvd.url)
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
