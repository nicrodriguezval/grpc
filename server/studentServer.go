package server

import (
	"context"
	"github.com/nicrodriguezval/grpc/models"
	"github.com/nicrodriguezval/grpc/repository"
	"github.com/nicrodriguezval/grpc/studentpb"
)

type StudentServer struct {
	repo repository.Repository
	studentpb.UnimplementedStudentServiceServer
}

func NewStudentServer(repo repository.Repository) *StudentServer {
	return &StudentServer{repo: repo}
}

func (s *StudentServer) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	student, err := s.repo.GetStudent(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &studentpb.Student{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func (s *StudentServer) CreateStudent(ctx context.Context, req *studentpb.Student) (*studentpb.CreateStudentResponse, error) {
	student := &models.Student{
		Id:   req.GetId(),
		Name: req.GetName(),
		Age:  req.GetAge(),
	}

	err := s.repo.CreateStudent(ctx, student)
	if err != nil {
		return nil, err
	}

	return &studentpb.CreateStudentResponse{
		Id: student.Id,
	}, nil
}
