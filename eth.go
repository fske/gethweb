package gethweb

import (
	"fmt"
)

type transactionParam struct {
	From     string
	To       string
	Gas      int
	GasPrice int
	Value    int
	Data     string
	Nonce    string
}

type contractParam struct {
	Name string
	Type string
	Indexed bool
}

type abi struct {
	Name     string
	Type     string
	Constant bool
	Inputs   []contractParam
	Outputs  []contractParam
}

type contract struct {
	AbiList []abi
	Address string
}

func (c *contract) At(address string) {
	c.Address = address
}

func (c contract) SendTransaction(method string, inputs []contractParam) {
	post("send_transaction", postRqst{})
}

type eth struct {}

func (e eth) Accounts() ([]string, error) {
	postData := postRqst{
		"eth_accounts",
		[]string{},
		getRandomInt(),
		"2.0",
	}
	postResult, err := post(hp.url, postData)
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

func (e eth) SendTransaction(txParam transactionParam) (string, error) {
	txParam.Gas = dec2Hex(txParam.Gas)
	tx.Param.GasPrice = dec2Hex(txParam.GasPrice)
	txParam.Value = dec2Hex(txParam.Value)

	params := []string{struct2LowerJson(txParam)}
	postData := postRqst{
		"eth_sendTransaction",
		params,
		getRandomInt(),
		"2.0",
	}
	postResult, err := post(hp.url, postData)
	return postResult.Result.(string), err
}

func (e eth) Contract(abiList []abi) contract {
	return contract{
		AbiList: abiList,
	}
} 


