package repository

import (
	"context"
	"github.com/nicrodriguezval/grpc/models"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	CreateStudent(ctx context.Context, student *models.Student) error
	GetTest(ctx context.Context, id string) (*models.Test, error)
	CreateTest(ctx context.Context, test *models.Test) error
}
