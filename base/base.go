package base

import (
	"CommunicationsModule/data"
	"fmt"
	"os"
)

var HttpPort = ":8080"
var Message = data.Message{}
var GrpcRemote = ""
var GrpcRemotePort = ""
var GrpcPort = ""
var IsGrpcServer = false



func init() {
	fmt.Println("Setting up Env")
	setUpPort()
	setUpGrpc()

}


func setUpPort() {
	port := os.Getenv("PORT")
	if port != "" {
		HttpPort = ":" + port
	}
}

func setUpGrpc() {
	GrpcRemote = os.Getenv("GRPC_REMOTE")
	GrpcPort = os.Getenv("GRPC_REMOTE_PORT")
	GrpcPort = os.Getenv("GRPC_PORT")

	if GrpcPort != "" {
		GrpcPort = ":" + GrpcPort
		IsGrpcServer = true
	}
}