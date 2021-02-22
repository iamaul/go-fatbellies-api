package utils

type ResponseJSON struct {
	Code    int         `json:"code,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Success bool        `json:"success,omitempty"`
}
