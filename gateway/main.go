package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/wodm8/backend-go/commons"
	pb "github.com/wodm8/backend-go/commons/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddress         = commons.EnvString("HTTP_ADDRESS", ":8080")
	exerciseServiceAddr = "localhost:2000"
)

func main() {
	conn, err := grpc.Dial(exerciseServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect to exercise service: ", err)
	}
	defer conn.Close()

	log.Printf("Dialing exercise service on %s", exerciseServiceAddr)

	c := pb.NewExerciseServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(c)

	handler.registerRoutes(mux)

	log.Printf("Starting server on %s", httpAddress)

	if err := http.ListenAndServe(httpAddress, mux); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
