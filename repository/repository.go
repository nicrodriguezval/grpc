package repository

import (
	"context"
	"github.com/nicrodriguezval/grpc/models"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return impl.GetStudent(ctx, id)
}

func SetStudent(ctx context.Context, student *models.Student) error {
	return impl.SetStudent(ctx, student)
}
