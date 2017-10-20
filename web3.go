package gethweb

import (
	"fmt"
)

type httpProvider struct {
	url          string
	timeOut      int64
	authUserName string
	authPassword string
}

// HttpProvider : initialize http config
func HttpProvider(url string, others ...interface{}) httpProvider {
	timeOut := int64(10)
	userName := ""
	password := ""
	for k, v := range others {
		switch k { 
			case 0:
				timeOut, _ = v.(int64)
			case 1:
				userName, _ = v.(string)
			case 2:
				password, _ = v.(string)
			default:
		}
	}
		
	return httpProvider {
		url:          url,
		timeOut:      timeOut,
		authUserName: userName,
		authPassword: password,
	}
}

// Web3 : module entry
type Web3 struct {
	Version      version
	Eth          eth
}

// NewWeb3 : return a web3 instance
func NewWeb3(provider httpProvider) *Web3 {
	fmt.Println("instanced")
	w := &Web3{
		Version: version{
			httpPrvd: provider,
		},
		Eth: eth{
			httpPrvd: provider,
		},
	}
	return w
}

func (w Web3) Sha3(input string) string {
	return sha3Keccak256(input)
}
