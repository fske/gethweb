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
	id      int
    jsonrpc string
    result  interface{}
	error   interface{}
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
		resp = postResp{}
		return resp, err
	}
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("read error:", err)
	}
	var rs interface{}
	err = json.Unmarshal(result, &rs)
	fmt.Println("rs:", rs)
	return rs, nil
}