package lib

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int    `json:"-"`
	IsError bool   `json:"-"`
	Status  string `json:"status"`

	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (res *Response) Send(w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")

	if res.IsError {
		w.WriteHeader(res.Code)
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		w.Write([]byte("Error when encode response"))
	}
}

func NewRes(data interface{}, code int, isError bool) *Response {

	if isError {
		return &Response{
			Code:    code,
			Status:  getStatus(code),
			IsError: isError,
			Message: data,
		}

	}
	return &Response{
		Code:    code,
		Status:  getStatus(code),
		IsError: isError,
		Message: "Success",
		Data:    data,
	}
}

func getStatus(status int) string {
	var desc string
	switch status {
	case 200:
		desc = "Success"
	case 201:
		desc = "Created"
	case 400:
		desc = "Bad Request"
	case 401:
		desc = "Unauthorized"
	case 404:
		desc = "Not Found"
	case 500:
		desc = "Internal Server Error"
	case 501:
		desc = "Bad Gateway"
	case 304:
		desc = "Not Modified"
	default:
		desc = ""
	}

	return desc
}
