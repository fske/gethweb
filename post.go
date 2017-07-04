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
	ID      int
    Jsonrpc string
}

type postResp struct{
	ID      int
    Jsonrpc string
    Result  interface{}
	Error   interface{}
}

func post(url string, data postRqst) (resp postResp, err error) {
	contentType := "application/json"
	dataString := struct2LowerJson(data)
	body := bytes.NewBuffer([]byte(dataString))

	response, err := http.Post(url, contentType, body)
	if err != nil {
		fmt.Println("post error:", err)
		return resp, err
	}
	defer response.Body.Close()

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