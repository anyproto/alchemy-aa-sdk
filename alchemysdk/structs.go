package alchemysdk

type EntryPointAddress interface{}

type UserOperation struct {
	Sender               string `json:"sender,omitempty"`
	Nonce                string `json:"nonce,omitempty"`
	InitCode             string `json:"initCode,omitempty"`
	CallData             string `json:"callData,omitempty"`
	Signature            string `json:"signature,omitempty"`
	CallGasLimit         string `json:"callGasLimit,omitempty"`
	VerificationGasLimit string `json:"verificationGasLimit,omitempty"`
	PreVerificationGas   string `json:"preVerificationGas,omitempty"`
	MaxFeePerGas         string `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas,omitempty"`
	PaymasterAndData     string `json:"paymasterAndData,omitempty"`
}

type GasAndPaymentStruct struct {
	PolicyID       string        `json:"policyId"`
	EntryPoint     string        `json:"entryPoint"`
	UserOperation  UserOperation `json:"userOperation"`
	DummySignature string        `json:"dummySignature"`
}

type JSONRPCRequestGasAndPaymaster struct {
	ID      int                   `json:"id"`
	JSONRPC string                `json:"jsonrpc"`
	Method  string                `json:"method"`
	Params  []GasAndPaymentStruct `json:"params"`
}

type JSONRPCResponseGasAndPaymaster struct {
	ID      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`

	Error struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"error,omitempty"`

	Result struct {
		PreVerificationGas   string `json:"preVerificationGas,omitempty"`
		CallGasLimit         string `json:"callGasLimit,omitempty"`
		VerificationGasLimit string `json:"verificationGasLimit,omitempty"`
		PaymasterAndData     string `json:"paymasterAndData,omitempty"`
		MaxFeePerGas         string `json:"maxFeePerGas,omitempty"`
		MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas,omitempty"`
	} `json:"result"`
}

type JSONRPCRequest struct {
	ID      int             `json:"id"`
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  []UserOperation `json:"params"`
}

type JSONRPCResponseUserOpHash struct {
	ID      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`

	Error struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"error,omitempty"`

	Result string `json:"result"`
}

type JSONRPCResponseGetOp struct {
	ID      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`

	Error struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"error,omitempty"`

	Result struct {
		UserOpHash string `json:"userOpHash,omitempty"`
		Success    bool   `json:"success,omitempty"`
	} `json:"result"`
}

type JSONRPCResponseGetUserOpByHash struct {
	ID      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`

	Error struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"error,omitempty"`

	Result struct {
		UserOperation   string `json:"userOperation,omitempty"`
		BlockNumber     string `json:"blockNumber,omitempty"`
		BlockHash       string `json:"blockNumber,omitempty"`
		TransactionHash string `json:"blockNumber,omitempty"`
	} `json:"result"`
}

type JSONRPCRequestGetUserOperationReceipt struct {
	ID      int      `json:"id"`
	JSONRPC string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Hashes  []string `json:"params"`
}
