package server

import (
	"context"
	"github.com/nicrodriguezval/grpc/models"
	testpb2 "github.com/nicrodriguezval/grpc/protos/testpb"
	"github.com/nicrodriguezval/grpc/repository"
)

type TestServer struct {
	repo repository.Repository
	testpb2.UnimplementedTestServiceServer
}

func NewTestServer(repo repository.Repository) *TestServer {
	return &TestServer{repo: repo}
}

func (s *TestServer) GetTest(ctx context.Context, req *testpb2.GetTestRequest) (*testpb2.Test, error) {
	test, err := s.repo.GetTest(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &testpb2.Test{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) CreateTest(ctx context.Context, req *testpb2.Test) (*testpb2.CreateTestResponse, error) {
	test := &models.Test{
		Id:   req.GetId(),
		Name: req.GetName(),
	}

	if err := s.repo.CreateTest(ctx, test); err != nil {
		return nil, err
	}

	return &testpb2.CreateTestResponse{
		Id: test.Id,
	}, nil
}
