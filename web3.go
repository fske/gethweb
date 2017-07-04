package gethweb

import (
	"fmt"
)

type httpProvider struct {
	url          string
	timeOut      int
	authUsername string
	authPassword string
}

// HttpProvider : initialize http config
func HttpProvider(url string, timeOut int, username string, password string) httpProvider {
	return httpProvider {
		url:          url,
		timeOut:      timeOut,
		authUsername: username,
		authPassword: password,
	}
}

var hp httpProvider

// Web3 : module entry
type Web3 struct {
	Version      version
	Provider     httpProvider
	Eth          eth
}

// NewWeb3 : return a web3 instance
func NewWeb3(provider httpProvider) *Web3 {
	hp = provider
	fmt.Println(hp)
	w := &Web3{
		Version:  version{},
		Provider: provider,
		Eth:      eth{},
	}
	return w
}
