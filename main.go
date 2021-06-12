package main

import (
	"CommunicationsModule/base"
	"CommunicationsModule/data"
	"CommunicationsModule/service"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func getTextHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, base.Message.ReadMessage())
}

func resetTextHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	base.Message.ResetMessage()
	fmt.Fprintf(w, base.Message.ReadMessage())
}

func addTextHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	payload := ""
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err == nil {
		payload = string(body)
	}

	base.Message.AddMessage(payload)
	fmt.Fprintf(w, base.Message.ReadMessage())
}

func main() {

	if base.IsGrpcServer {
		go serveGrpc()
		serveHttp()
	} else {
		// Set Up Initial Messages / Wait for Server
		for !service.GetExistingMessages() {
			time.Sleep(5 * time.Second)
		}

		// Update Messages
		for {
			service.SendMessage()
			time.Sleep(5 * time.Second)
		}
	}
}


func serveHttp() {
	router := httprouter.New()

	router.GET("/", getTextHandler)
	//router.PUT("/", addTextHandler)
	//router.POST("/", addTextHandler)
	router.DELETE("/", resetTextHandler)

	fmt.Printf("Starting HTTP Server at port %s \n", base.HttpPort)
	l, err := net.Listen("tcp4", base.HttpPort)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(l, router))
}

func serveGrpc() {
	fmt.Printf("Starting gRPC Server at port %s \n", base.GrpcPort)
	lis, err := net.Listen("tcp", base.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := service.Server{}

	grpcServer := grpc.NewServer()

	data.RegisterMessageServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}



