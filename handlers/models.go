package handlers

type apiError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
