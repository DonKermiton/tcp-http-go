package tcp_http_go

import (
	"bufio"
	"io"
	"strings"
)

const (
	GET     = "GET"
	HEAD    = "HEAD"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	DELETE  = "DELETE"
	OPTIONS = "OPTIONS"
)

type headers struct {
	headers map[string]string
}

func newHeaders() *headers {
	return &headers{
		headers: make(map[string]string),
	}
}

func (h *headers) decode(reader io.Reader) error {
	headerScanner := bufio.NewScanner(reader)

	for headerScanner.Scan() {
		rawHeader := headerScanner.Text()
		header := strings.SplitN(rawHeader, ":", 2)

		if len(header) == 2 {
			h.headers[header[0]] = strings.TrimSpace(header[1])
		}

		if rawHeader == "" {
			break
		}
	}

	return nil
}
