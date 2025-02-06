package main

import (
	"errors"
	"regexp"
	"strings"
)

type Method string
type httpVersion string

const (
	SPACE = " "
	CLRF  = "\r\n"
)
const (
	GET  Method = "GET"
	POST Method = "POST"
)

const (
	HTTP1   httpVersion = "HTTP/1.0"
	HTTP1_1 httpVersion = "HTTP/1.1"
	HTTP2   httpVersion = "HTTP/2.0"
	HTTP3   httpVersion = "HTTP/3.0"
)

type httpRequest struct {
	Method      Method
	URI         string
	HttpVersion httpVersion
}

func httpParser(b []byte) (*httpRequest, error) {
	request := strings.Split(strings.TrimSpace(string(b)), "\n")
	httpReq, err := parseLine(strings.TrimSpace(request[0]))
	if err != nil {
		return nil, err
	}
	return httpReq, nil

}
func parseLine(s string) (*httpRequest, error) {
	re := regexp.MustCompile(`(?m)\/([a-zA-Z0-9-._~]:?){0,20}([a-zA-Z0-9!*$$;:@&=+\$,\/\?$$$$\-\.~%_]+:?)?`)
	httpObject := &httpRequest{}
	splitted := strings.Split(s, SPACE)

	if len(splitted) != 3 {
		return nil, errors.New("Invalid Format")
	}
	switch splitted[0] {
	case "GET":
		httpObject.Method = GET
	case "POST":
		httpObject.Method = POST
	default:
		return nil, errors.New("Malformed Http request")
	}
	if !re.MatchString(splitted[1]) {
		return nil, errors.New("Malformed URI request")
	}
	httpObject.URI = splitted[1]
	switch splitted[2] {
	case "HTTP/1.0", "HTTP/1":
		httpObject.HttpVersion = HTTP1
	case "HTTP/1.1":
		httpObject.HttpVersion = HTTP1_1
	case "HTTP/2", "HTTP/2.0":
		httpObject.HttpVersion = HTTP2
	case "HTTP/3.0", "HTTP/3":
		httpObject.HttpVersion = HTTP3
	default:
		return nil, errors.New("Malfromed http version")
	}
	return httpObject, nil
}
