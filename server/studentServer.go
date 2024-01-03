package server

import (
	"context"
	"github.com/nicrodriguezval/grpc/models"
	studentpb2 "github.com/nicrodriguezval/grpc/protos/studentpb"
	"github.com/nicrodriguezval/grpc/repository"
)

type StudentServer struct {
	repo repository.Repository
	studentpb2.UnimplementedStudentServiceServer
}

func NewStudentServer(repo repository.Repository) *StudentServer {
	return &StudentServer{repo: repo}
}

func (s *StudentServer) GetStudent(ctx context.Context, req *studentpb2.GetStudentRequest) (*studentpb2.Student, error) {
	student, err := s.repo.GetStudent(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &studentpb2.Student{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func (s *StudentServer) CreateStudent(ctx context.Context, req *studentpb2.Student) (*studentpb2.CreateStudentResponse, error) {
	student := &models.Student{
		Id:   req.GetId(),
		Name: req.GetName(),
		Age:  req.GetAge(),
	}

	err := s.repo.CreateStudent(ctx, student)
	if err != nil {
		return nil, err
	}

	return &studentpb2.CreateStudentResponse{
		Id: student.Id,
	}, nil
}
