package service

import (
"log"
	"CommunicationsModule/base"
	"CommunicationsModule/data"
"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *data.MyMessage) (*MyMessage, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &MyMessage{Body: "Hello From the Server!"}, nil
}
