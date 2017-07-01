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

func HttpProvider(url string, timeOut int, username string, password string) httpProvider {
	return httpProvider {
		url:          url,
		timeOut:      timeOut,
		authUsername: username,
		authPassword: password,
	}
}

var hp httpProvider

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
