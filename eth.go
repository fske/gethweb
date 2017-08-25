package gethweb

import (
	"fmt"
)

type eth struct {
	httpPrvd     httpProvider
}

func (e eth) Accounts() ([]string, error) {
	postData := postRqst{
		"eth_accounts",
		[]string{},
		getRandomInt(),
		"2.0",
	}
	postResult, err := post(e.httpPrvd.url, postData)
	if err != nil {
		fmt.Println(err)
		return []string{}, err
	}
	//fmt.Println(postResult.Result)
	var accounts []string
	for _, result := range postResult.Result.([]interface{}) {
		account, ok := result.(string)
		if ok == true {
			accounts = append(accounts, account)
		}
	}
	return accounts, nil
}
func (e eth) SendTransaction(in map[string]interface{}) (string, error) {
	return sendTransaction(e.httpPrvd.url, in)
}

func (e eth) Contract(abiList []abi) contract {
	return contract{
		AbiList: abiList,
		httpPrvd: e.httpPrvd
	}
} 

type abi struct {
	Name     string
	Type     string
	Constant bool
	Inputs   []contractParam
	Outputs  []contractParam
}

type contract struct {
	httpPrvd        httpProvider
	AbiList         []abi
	TransactionHash string
	Address         string
}

func (c *contract) New(in map[string]interface{}) (*contract, error) {
	resp, err := sendTransaction(c.httpPrvd.url, in)
	return resp, err
}

func (c *contract) At(address string) (*contract) {
	c.Address = address
	return c
}

func (c contract) SendTransaction(method string, inputs []contractParam) {
	post("send_transaction", postRqst{})
}


