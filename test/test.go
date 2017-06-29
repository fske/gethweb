package main

import (
  "fmt"

  "gethweb"
)

func main() {
  var x gethweb.Web3
  x.Provider.HttpProvider("ddd", 12, "user", "pass")
  fmt.Printf("%v\n", x.Provider)
  fmt.Println(x.Version.Api())
}
