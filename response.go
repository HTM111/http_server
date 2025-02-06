package main

import (
	"bytes"
	"fmt"
	"net"
)

type ResponseWrite struct {
	statusCode int
	Con        net.Conn
	headers    map[string]string
}

func WriteResponse(statusCode int, connection net.Conn) {
	r := ResponseWrite{Con: connection, headers: map[string]string{}}
	r.WriteHeader(statusCode)
	r.Write([]byte(""))
}
func (r *ResponseWrite) AddHeader(key string, value string) {
	r.headers[key] = value
}
func (r *ResponseWrite) WriteHeader(i int) {
	r.statusCode = i
}
func (r *ResponseWrite) Write(b []byte) {

	var formatedHeaders bytes.Buffer
	var textMessage = "Found"
	if len(r.headers) == 0 {
		r.headers["Content-Type"] = "text/html"
	}
	if r.statusCode == 0 {
		r.statusCode = 200
	}
	for key, value := range r.headers {
		formatedHeaders.Write([]byte(fmt.Sprintf("%s:%s\n", key, value)))
	}
	if value, exist := httpStatusCodes[r.statusCode]; exist {
		textMessage = value
	}
	formatedRequest := fmt.Sprintf("HTTP/1.0 %d  %s\n%s\n%s", r.statusCode, textMessage, formatedHeaders.String(), b)
	r.Con.Write([]byte(formatedRequest))
	defer formatedHeaders.Reset()
}
