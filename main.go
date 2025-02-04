package main

import (
	"log"
	"net"

	"github.com/Ukkenjijo/email-service/proto"
	"github.com/Ukkenjijo/email-service/service"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterEmailServiceServer(grpcServer, &service.EmailService{})

	log.Println("Email Service is running on port 50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
