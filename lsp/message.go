package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`

	// Params ...
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id,omitempty"`

	// Result
	// Error
}

type Notification struct {
	RRC    string `json:"jsonrpc"`
	Method string `json:"method"`
}
