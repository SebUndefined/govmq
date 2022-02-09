package govmq

type Response interface {
}

type OKResponse struct {
	Result    string      `json:"result"`
	Modifiers interface{} `json:"modifiers"`
}

// ErrorResponse an error response that satisfy VMQ Broker
type ErrorResponse struct {
	Result struct {
		Error string `json:"error"`
	} `json:"result"`
}

func NewErrorResponse(result struct {
	Error string `json:"error"`
}) *ErrorResponse {
	return &ErrorResponse{Result: result}
}

type Builder interface {
	Build() Response
}
