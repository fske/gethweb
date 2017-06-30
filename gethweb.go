package gethweb

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type httpProvider struct {
	url          string
	timeOut      int
	authUsername string
	authPassword string
}

func HttpProvider(url string, timeOut int, username string, password string) httpProvider {
	return httpProvider {
		url:          url,
		timeOut:      timeOut,
		authUsername: username,
		authPassword: password,
	}
}

var hp httpProvider

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

type eth struct {}

func (e eth) Accounts() ([]string) {
	postData := postRqst{"eth_accounts", []string{}, 1, "2.0"}
	postResult, err := post(hp.url, postData)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	//fmt.Println(postResult.Result)
	var accounts []string
	for _, result := range postResult.Result.([]interface{}) {
		account, ok := result.(string)
		if ok == true {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

type Web3 struct {
	Version      version
	Provider     httpProvider
	Eth          eth
}

func NewWeb3(provider httpProvider) *Web3 {
	hp = provider
	w := &Web3{
		Version:  version{},
		Provider: provider,
		Eth:      eth{},
	}
	return w
}
