package main

import (
	"context"
	"log"
	"net"

	_ "github.com/joho/godotenv/autoload"
	"github.com/wodm8/backend-go/commons"
	"google.golang.org/grpc"
)

var (
	grpcAddress = commons.EnvString("GRPC_ADDRESS", "localhost:3000")
)

func main() {
	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}
	defer l.Close()

	store := NewStore()
	svc := NewService(store)

	NewGrpcHandler(grpcServer)

	svc.CreateExercise(context.Background())

	log.Printf("Starting grpc server on %s", grpcAddress)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}
