package main

import (
	"errors"
	"net/http"

	"github.com/wodm8/backend-go/commons"
	pb "github.com/wodm8/backend-go/commons/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	client pb.ExerciseServiceClient
}

func NewHandler(c pb.ExerciseServiceClient) *handler {
	return &handler{client: c}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	// Register route for create exercise with post method
	mux.HandleFunc("POST /api/exercises", h.CreateExercise)
}

func (h *handler) CreateExercise(w http.ResponseWriter, r *http.Request) {
	var exerciseRequest pb.CreateExcerciseRequest
	if err := commons.ReadJSON(r, &exerciseRequest); err != nil {
		commons.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateExercise(&exerciseRequest); err != nil {
		commons.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	e, err := h.client.CreateExercise(r.Context(), &pb.CreateExcerciseRequest{
		Name:     exerciseRequest.Name,
		Category: exerciseRequest.Category,
	})

	rStatus := status.Convert(err)
	if rStatus != nil {
		if rStatus.Code() != codes.InvalidArgument {
			commons.WriteError(w, http.StatusBadRequest, rStatus.Message())
			return
		}
	}

	commons.WriteJSON(w, http.StatusCreated, e)
}

func validateExercise(exerciseRequest *pb.CreateExcerciseRequest) error {
	if exerciseRequest.Name == "" {
		return errors.New("name is required")
	}
	if exerciseRequest.Category == "" {
		return errors.New("category is required")
	}
	return nil
}
