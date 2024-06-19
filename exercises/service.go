package main

import "context"

type Service struct {
	store ExerciseStore
}

func NewService(store ExerciseStore) *Service {
	return &Service{store: store}
}

func (s *Service) CreateExercise(ctx context.Context) error {
	return nil
}
