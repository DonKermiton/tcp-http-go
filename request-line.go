package tcp_http_go

import (
	"fmt"
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

func newRequestLine(line string) *requestLine {
	return &requestLine{
		rawRequestLine: line,
	}
}

func (r *requestLine) decode() error {
	requestComponents := strings.Split(r.rawRequestLine, " ")

	if len(requestComponents) != 3 {
		return fmt.Errorf("provided wrong number of args in http")
	}

	r.method = requestComponents[0]
	r.path = requestComponents[1]
	r.httpVersion = requestComponents[2]

	return nil
}
