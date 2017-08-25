package gethweb

type sendTransactionInput struct {
	From     string  `json: "from" required: "true"`
	To       string  `json: "to"`
	Gas      int     `json: "gas" hex: "true"`
	GasPrice int     `json: "gasPrice" hex: "true"`
	Value    int     `json: "value" hex: "true"`
	Data     string  `json: "data"`
	Nonce    string  `json: "nonce" hex: "true"`
}

type contractParam struct {
	Name string
	Type string
	Indexed bool
}

