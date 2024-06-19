package main

import (
	"context"

	pb "github.com/wodm8/backend-go/commons/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedExerciseServiceServer
}

func NewGrpcHandler(grpcServer *grpc.Server) *grpcHandler {
	handler := &grpcHandler{}
	pb.RegisterExerciseServiceServer(grpcServer, handler)
	return handler
}

func (h *grpcHandler) CreateExercise(ctx context.Context, p *pb.CreateExcerciseRequest) (*pb.Excercise, error) {
	exercise := &pb.Excercise{
		ID:       "11223344-5566-7788-9900-aabbccddeeff",
		Name:     p.Name,
		Category: p.Category,
	}
	return exercise, nil
}
