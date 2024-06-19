package main

import "context"

type ExerciseService interface {
	CreateExercise(context.Context) error
}

type ExerciseStore interface {
	Create(context.Context) error
}
