package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"time"
)

type Handler func(w httpWriter)
type Server struct {
	routing     map[string]Handler
	Adress      string
	bytesWriter []byte
}

func NewServer(Address string) *Server {
	return &Server{
		routing: make(map[string]Handler, 0),
		Adress:  Address,
	}
}

func (s *Server) ListenAndServe() error {
	address := fmt.Sprintf("%s", s.Adress)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error Listener.Accept() : ", err)
			continue
		}
		go s.HandleConnection(conn)
		defer conn.Close()
	}
}
func (s *Server) Route(path string, handler Handler) {
	s.routing[path] = handler

}
func (s *Server) HandleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))

		buffer := make([]byte, 1024)
		var req bytes.Buffer

		for {
			n, err := conn.Read(buffer)
			if err != nil {
				log.Println("Unable to read connection", err)
				return
			}
			req.Write(buffer[:n])
			if req.Len() > 1 {
				break
			}

		}
		httpResponse, err := httpParser(buffer[:100])
		if err != nil {
			WriteResponse(400, conn)
			return
		}

		if s.routing[httpResponse.URI] != nil {
			handler := s.routing[httpResponse.URI]
			handler(&ResponseWrite{Con: conn, headers: map[string]string{}})
			break
		}

		WriteResponse(404, conn)
		break
	}
	conn.Close()
}
