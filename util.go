package gethweb

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

func dec2Hex(decimal int64) (hex string) {
	return fmt.Sprintf("0x%x", decimal)
}

func struct2LowerJson(strc interface{}) (strg string) {
	b, err := json.Marshal(strc)
	if err != nil {
		return strg
	}
	strg = strings.ToLower(string(b))
	return strg
}

func map2LowerJson(m map[string]string) (strg string) {
	b, err := json.Marshal(m)
	if err != nil {
		return strg
	}
	strg = strings.ToLower(string(b))
	return strg
}

func getRandomInt() (int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000)
}

func paramUniform(paramSet map[string]interface{}, proto interface{}) (map[string]string, error) {
	uniParamSet := map[string]string{}
	protoValue := reflect.ValueOf(proto).Elem()
	for i, len := 0, protoValue.NumField(); i < len; i++ {
		protoDesc := protoValue.Type().Field(i)
		protoParamName := protoDesc.Name
		paramValue, ok := paramSet[protoParamName]
		fmt.Println(ok)
		if ok != true {
			if protoDesc.Tag.Get("required") == "true" {
				return map[string]string{}, errors.New("missing field: " + protoParamName)
			}
			continue
		}
		if protoDesc.Type != reflect.TypeOf(paramValue) {
			fmt.Println(reflect.TypeOf(paramValue).String())
			return map[string]string{}, errors.New("field type error: " + protoParamName)
		}
		switch paramValue.(type) {
			case int64:
				if protoDesc.Tag.Get("hex") != "true" {
					uniParamSet[protoParamName] = fmt.Sprintf("%d", paramValue)
				} else {
					uniParamSet[protoParamName] = fmt.Sprintf("%d", paramValue)
				}
			default:
				uniParamSet[protoParamName] = fmt.Sprintf("%s", paramValue)
		}
	}
	return uniParamSet, nil
}

