package main

import (
	"fmt"

	"github.com/fske/gethweb"
)

func main() {
	x := gethweb.NewWeb3(gethweb.HttpProvider("http://123.59.80.44:8540", 12, "user", "pass"))
	//fmt.Printf("%v\n", x.Provider)
	//fmt.Println(x.Version.Node())

	//params := []string{}
	results, _ := x.Eth.Accounts()
	fmt.Println(results)

	result := x.Sha3("some data to hash")
	fmt.Println(result)
}
