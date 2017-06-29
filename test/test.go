package main

import (
  "fmt"

  "gethweb"
)

func main() {
  var x gethweb.NewWeb3
  x.Provider.HttpProvider("www.baidu.com", 12, "user", "pass")
  fmt.Printf("%v\n", x.Provider)
  fmt.Println(x.Version.Node())
}
