package gethweb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type postRqst struct{
	Method  string
	Params  []string
	Id      int
    Jsonrpc string
}

type postResp struct{
	Id      int
    Jsonrpc string
    Result  interface{}
	Error   interface{}
}

func post(url string, data postRqst) (resp postResp, err error) {
	contentType := "application/json"
	b, err := json.Marshal(data)
	b_lower := strings.ToLower(string(b))
	body := bytes.NewBuffer([]byte(b_lower))

	response, err := http.Post(url, contentType, body)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("post error:", err)
		return resp, err
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("read error:", err)
		return resp, err
	}
	err = json.Unmarshal(responseBody, &resp)
	if err != nil {
		return resp, err
	}
	//fmt.Println("resp:", resp)
	return resp, nil
}