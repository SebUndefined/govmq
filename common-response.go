package govmq

type OKResponse struct {
	Result    string   `json:"result"`
	Modifiers Modifier `json:"modifiers"`
}

func NewOKResponse(modifiers Modifier) *OKResponse {
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

func NewErrorResponse(m string) *ErrorResponse {
	return &ErrorResponse{struct {
		Error string `json:"error"`
	}{Error: m}}
}
