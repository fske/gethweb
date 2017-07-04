package gethweb

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func dec2Hex(decimal int) (hex string) {
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

func getRandomInt() (int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000)
}