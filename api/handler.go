package api

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

// SayHello prints out the message recieved from the client and relays a message back
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Recieved Message from client %s", in.Greeting)
	return &PingMessage{Greeting: "Pong"}, nil
}
