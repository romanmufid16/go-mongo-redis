package model

import "strings"

type WebResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  interface{} `json:"result,omitempty"`
}

func BuildResponse(message string, result interface{}) *WebResponse {
	res := &WebResponse{
		Success: true,
		Message: message,
		Result:  result,
	}
	return res
}

type ErrorResponse struct {
	Success bool        `json:"success"`
	Errors  interface{} `json:"errors"`
}

func BuildErrorResponse(errors string) *ErrorResponse {
	splitErrors := strings.Split(errors, "\n")
	res := &ErrorResponse{
		Success: false,
		Errors:  splitErrors,
	}
	return res
}
