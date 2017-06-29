package gethweb

import (
  "io/ioutil"
  "net/http"
)

type httpProvider struct {
  url          string
  timeOut      int
  authUsername string
  authPassword string
}

func (hp *httpProvider) HttpProvider(url string, timeOut int, username string, password string) {
  *hp = httpProvider {
    url          : url,
    timeOut      : timeOut,
    authUsername : username,
    authPassword : password,
  }
}

type version struct {}

func (v version) Api() (string) {
  return "0.0.1"
}

func (v version) Node() (string) {
  var result string
  resp, err := http.Get(web3.Provider.url)
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

//func (e eth) GetAccounts() ([]string) {
//  var results []string
//  resp, err := http.Post()
//}

type Web3 struct {
  Version      version
  Provider     httpProvider
  Eth          eth
}

var web3 Web3

func NewWeb3() Web3 {
  return web3
}
