package govmq

type Response interface {
}

type OKResponse struct {
	Result    string      `json:"result"`
	Modifiers interface{} `json:"modifiers"`
}

func NewOKResponse(modifiers interface{}) *OKResponse {
	return &OKResponse{Result: "ok", Modifiers: modifiers}
}

type EmptyResponse struct {
}

func NewEmptyResponse() *EmptyResponse {
	return &EmptyResponse{}
}

type NextResponse struct {
	Result string `json:"result"`
}

func NewNextResponse() *NextResponse {
	return &NextResponse{Result: "next"}
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

type responseBuilderAction func(Response)
