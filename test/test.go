package main

import (
	"fmt"

	"gethweb"
)

func main() {
	x := gethweb.NewWeb3(gethweb.HttpProvider("http://123.59.80.44:8540", 12, "user", "pass"))
	fmt.Printf("%v\n", x.Provider)
	//fmt.Println(x.Version.Node())

	//params := []string{}
	//data := gethweb.PostRqst{"eth_accounts", params, 1, "2.0"}
	result := x.Eth.Accounts()
	fmt.Println(result)
}
