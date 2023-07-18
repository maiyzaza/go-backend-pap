package handler

type Wrapper struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}

type MexceptionError struct {
	Code    string `json:"code"`
	Title   string `json:"title"`
	Message string `json:"message"`
}
