package tcp_http_go

import (
	"bufio"
	"io"
	"strings"
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
		h.headers[header[0]] = header[1]

		if rawHeader == "" {
			break
		}
	}

	return nil
}
