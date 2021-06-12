package service

import (
	"CommunicationsModule/base"
	"CommunicationsModule/data"
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) SendMessage(ctx context.Context, in *data.MyMessage) (*data.MyMessage, error) {
	log.Printf("Receive message body from client: %s", in.Message)

	base.Message.SetMessage(in.Message)
	base.Message.AddMessage(RandomLoremIpsum())

	return &data.MyMessage{Message: base.Message.ReadMessage()}, nil
}

func (s *Server) GetMessage(ctx context.Context, in *data.Empty) (*data.MyMessage, error) {
	log.Printf("Setting Up Existing Messages: %s", base.Message.ReadMessage())
	return &data.MyMessage{Message: base.Message.ReadMessage()}, nil
}