package gethweb

import (
	"fmt"
)

type param struct {
	Name string
	Type string
	Indexed bool
}

type abi struct {
	Name     string
	Type     string
	Constant bool
	Inputs   []param
	Outputs  []param
}

type contract struct {
	AbiList []abi
}

func (c *contract) At(address string) {
	c.Address = address
}
func (c contract) SendTransaction(method string, inputs []param) {
	post("send_transaction", "")
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

func (e eth) Contract(abi_list []abi) contract {
	return contract{
		AbiList: abiList
	}
} 


