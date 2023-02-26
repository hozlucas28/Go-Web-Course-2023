package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status        int         `json:"status"`
	Data          interface{} `json:"data"`
	Message       string      `json:"message"`
	contentType   string
	responseWrite http.ResponseWriter
}

func CreateDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:        http.StatusOK,
		responseWrite: rw,
		contentType:   "application/json",
	}
}

func (r *Response) Send() {
	r.responseWrite.Header().Set("content-type", r.contentType)
	r.responseWrite.WriteHeader(r.Status)

	output, _ := json.Marshal(&r)
	fmt.Fprintln(r.responseWrite, string(output))
}

func SendData(rw http.ResponseWriter, data interface{}) {
	r := CreateDefaultResponse(rw)
	r.Data = data
	r.Send()
}

func (r *Response) NotFound() {
	r.Status = http.StatusNotFound
	r.Message = "Resource not found"
}

func SendNotFound(rw http.ResponseWriter) {
	r := CreateDefaultResponse(rw)
	r.NotFound()
	r.Send()
}

func (r *Response) UnprocessableEntity() {
	r.Status = http.StatusUnprocessableEntity
	r.Message = "Unprocessable entity not found"
}

func SendUnprocessableEntity(rw http.ResponseWriter) {
	r := CreateDefaultResponse(rw)
	r.UnprocessableEntity()
	r.Send()
}