package model

type SignUpResponse struct {
	StatusCode int32  `json:"statusCode"`
	Message    string `json:"message"`
}

type LoginResponse struct {
	StatusCode int32  `json:"statusCode"`
	Message    string `json:"message"`
	Token      string `json:"token"`
}
