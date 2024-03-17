package response

import (
	"net/http"
)

type (
	Response struct {
		Status  int            `json:"status"`
		Message string         `json:"message"`
		Errors  []CaptureError `json:"errors,omitempty"`
		Data    interface{}    `json:"data,omitempty"`
		Meta    interface{}    `json:"meta,omitempty"`
		Header  http.Header    `json:"header,omitempty"`
		Body    interface{}    `json:"body,omitempty"`
	}

	CaptureError struct {
		Details string `json:"details"`
		Message string `json:"message"`
	}
)

var (
	Text = http.StatusText

	MsgSuccess = "Success"
	MsgFailed  = "Failed"
)

func NewResponse(statusCode int, message string, data interface{}) Response {
	return Response{
		Status:  statusCode,
		Message: MsgSuccess,
		Data:    data,
	}
}

func NewResponseError(statusCode int, messageStatus string, details string) Response {
	return Response{
		Status:  statusCode,
		Message: messageStatus,
		Errors: []CaptureError{
			{
				Message: Text(statusCode),
				Details: details,
			},
		},
	}
}
