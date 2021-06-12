package service

import (
	"CommunicationsModule/base"
	"CommunicationsModule/data"
	"context"
	"google.golang.org/grpc"
	"log"
	"math/rand"
)

func GetExistingMessages() bool {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(base.GrpcRemoteUri, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()


	messageClient := data.NewMessageServiceClient(conn)
	response, err := messageClient.GetMessage(context.Background(), &data.Empty{
	})

	if err == nil {
		base.Message.SetMessage(response.GetMessage())
		return true
	}
	log.Printf("Error reading existing messages: %s", err.Error())
	return false
}

func SendMessage() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(base.GrpcRemoteUri, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()


	messageClient := data.NewMessageServiceClient(conn)

	message := base.Message.ReadMessage() + RandomLoremIpsum()

	response, err := messageClient.SendMessage(context.Background(), &data.MyMessage{
		Message: message,
	})

	log.Printf("Receive message body from server: %s", response.GetMessage())
	base.Message.SetMessage(response.GetMessage())
}


func RandomLoremIpsum() string {
	return LoremIpsumValues[rand.Intn(LoremIpsumLength)] + " "
}