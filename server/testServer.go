package server

import (
	"context"
	"github.com/nicrodriguezval/grpc/models"
	"github.com/nicrodriguezval/grpc/repository"
	"github.com/nicrodriguezval/grpc/testpb"
)

type TestServer struct {
	repo repository.Repository
	testpb.UnimplementedTestServiceServer
}

func NewTestServer(repo repository.Repository) *TestServer {
	return &TestServer{repo: repo}
}

func (s *TestServer) GetTest(ctx context.Context, req *testpb.GetTestRequest) (*testpb.Test, error) {
	test, err := s.repo.GetTest(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &testpb.Test{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) CreateTest(ctx context.Context, req *testpb.Test) (*testpb.CreateTestResponse, error) {
	test := &models.Test{
		Id:   req.GetId(),
		Name: req.GetName(),
	}

	if err := s.repo.CreateTest(ctx, test); err != nil {
		return nil, err
	}

	return &testpb.CreateTestResponse{
		Id: test.Id,
	}, nil
}
