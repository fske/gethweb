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

func (hp httpProvider) HttpProvider(url string, timeOut int, username string, password string) {
  hp = httpProvider {
    url          : url,
    timeOut      : timeOut,
    authUsername : username,
    authPassword : password,
  }
  fmt.Println(hp)
}

type version struct {}

func (v version) Api() (string) {
  return "0.0.1"
}

type Web3 struct {
  Version      version
  Provider     httpProvider
}