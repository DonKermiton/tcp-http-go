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
	requestComponents, err := splitRequestHeaderAndValidate(r.rawRequestLine)

	if err != nil {
		return err
	}

	if isMethodCorrect(requestComponents[0]) == false {
		return fmt.Errorf("invalid method")
	}

	r.method = requestComponents[0]
	r.path = requestComponents[1]

	switch len(requestComponents) {
	case 2:
		r.httpVersion = HTTP09
		break
	case 3:
		if isHttpVersionCorrect(requestComponents[2]) == false {
			return fmt.Errorf("invalid http protocol version")
		}

		r.httpVersion = requestComponents[2]
		break
	default:
		return fmt.Errorf("something went wrong with parsing http request line. Check is valid")
	}

	return nil
}

func splitRequestHeaderAndValidate(rawRequestLine string) ([]string, error) {
	requestComponents := strings.Split(rawRequestLine, " ")

	if len(requestComponents) != 2 && len(requestComponents) != 3 {
		return nil, fmt.Errorf("invalid number of components in request line")
	}

	return requestComponents, nil
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
