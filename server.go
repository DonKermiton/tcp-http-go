package tcp_http_go

import (
	"bufio"
	"fmt"
	"net"
)

type Router struct {
	port        uint32
	conn        net.Conn
	requestLine *requestLine
}

func NewRouter() *Router {
	return &Router{
		port: 4220,
	}

}

func (r *Router) Listen() error {
	server, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", r.port))

	if err != nil {
		return fmt.Errorf("error while starting a server: %w", err) // Wrap the error
	}
	for {
		conn, err := server.Accept()

		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		r.conn = conn
		go func() {
			err := r.handleConnection()
			if err != nil {

			}
		}()
	}

}

func (r *Router) handleConnection() error {
	reader := bufio.NewReader(r.conn)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	// request line
	requestLineStruct := newRequestLine(requestLine)
	err = requestLineStruct.decode()

	if err != nil {
		return err
	}

	r.requestLine = requestLineStruct

	// headers

	headersStruct := newHeaders()
	err = headersStruct.decode(reader)

	return nil
}
