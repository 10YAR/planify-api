package types

type HttpResponse struct {
	Status   int
	Message  string
	HttpCode int
	Error    error  `json:"Error,omitempty"`
	Token    string `json:"Token,omitempty"`
}
