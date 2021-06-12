package base

import (
	"CommunicationsModule/data"
	"fmt"
	"os"
)

var HttpPort = ":8080"
var Message = data.Message{}
var GrpcRemoteUri = ""
var GrpcPort = ""
var IsGrpcServer = false



func init() {
	fmt.Println("Setting up Env")
	setUpHttpPort()
	setUpGrpc()

}


func setUpHttpPort() {
	port := os.Getenv("PORT")
	if port != "" {
		HttpPort = ":" + port
	}
}

func setUpGrpc() {

	GrpcPort = os.Getenv("GRPC_PORT")

	if GrpcPort != "" {
		GrpcPort = ":" + GrpcPort
		IsGrpcServer = true
	}

	GrpcRemoteUri = os.Getenv("GRPC_REMOTE_URI")
}