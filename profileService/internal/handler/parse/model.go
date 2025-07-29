package parse

type ParseResponce struct {
	Message string `json:"message"`
	Error   error  `json:"error,omitempty"`
}
