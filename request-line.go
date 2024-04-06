package tcp_http_go

import (
	"fmt"
	"github.com/DonKermiton/tcp-http-go/utils"
	"strings"
)

type requestLine struct {
	//eg "GET /echo HTTP/1.1"
	rawRequestLine string
	//eg "GET"
	method string
	//eg "/echo"
	path string
	//eg "HTTP/1.1"
	httpVersion string
}

const (
	GET     = "GET"
	HEAD    = "HEAD"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	DELETE  = "DELETE"
	OPTIONS = "OPTIONS"
)

const (
	HTTP09 = "HTTP/0.9"
	HTTP1  = "HTTP/1.0"
	HTTP11 = "HTTP/1.1"
)

func newRequestLine(line string) *requestLine {
	return &requestLine{
		rawRequestLine: line,
	}
}

func (r *requestLine) decode() error {
	requestComponents := strings.Split(r.rawRequestLine, " ")

	numberOfRequestComponents := len(requestComponents)

	if numberOfRequestComponents < 2 || numberOfRequestComponents > 3 {
		return fmt.Errorf("provided wrong number of args in http")
	}

	validRequestHeader := isMethodCorrect(requestComponents[0]) && isHttpVersionCorrect(requestComponents[1])

	if !validRequestHeader {
		return fmt.Errorf("provided wrong protocol version or method type")
	}

	if numberOfRequestComponents == 2 {
		r.httpVersion = HTTP09
	}

	r.method = requestComponents[0]
	r.path = requestComponents[1]

	if numberOfRequestComponents == 2 {
		r.httpVersion = HTTP09
	} else if numberOfRequestComponents == 3 {
		r.httpVersion = requestComponents[2]
	}

	return nil
}

func isMethodCorrect(method string) bool {
	methods := []string{GET, HEAD, PATCH, POST, PUT, DELETE, OPTIONS}
	_, index := utils.Find(methods, func(s string) bool {
		return s == method
	})

	return index >= 0
}

func isHttpVersionCorrect(httpVersion string) bool {
	return httpVersion == HTTP1 || httpVersion == HTTP11
}
