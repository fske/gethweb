package gethweb

import (
	"fmt"
)

type transactionParams struct {
	From     string
	To       string
	Gas      int
	GasPrice int
	Value    int
	Data     string
	Nonce    string
}

type contractParams struct {
	Name string
	Type string
	Indexed bool
}

type abi struct {
	Name     string
	Type     string
	Constant bool
	Inputs   []contractParams
	Outputs  []contractParams
}

type contract struct {
	AbiList []abi
	Address string
}

func (c *contract) At(address string) {
	c.Address = address
}
func (c contract) SendTransaction(method string, inputs []contractParams) {
	post("send_transaction", postRqst{})
}
type eth struct {}

func (e eth) Accounts() ([]string) {
	postData := postRqst{
		"eth_accounts",
		[]string{},
		getRandomInt(),
		"2.0",
	}
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

func (e eth) SendTransaction(txParams transactionParams) (string) {
	`for key, value := range txParams {`
		if key == "Gas" || key == "GasPrice" || key == "Value" {
			txParams[key] = dec2Hex(value)
		}
	}
	params := []string{struct2LowerJson(txParams)}
	postData := postRqst{
		"eth_sendTransaction",
		params,
		getRandomInt(),
		"2.0",
	}
	postResult := post(hp.url, postData)
	return postResult.Result.(string)
}

func (e eth) Contract(abiList []abi) contract {
	return contract{
		AbiList: abiList,
	}
} 


