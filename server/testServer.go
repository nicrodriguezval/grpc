package server

import (
	"context"
	"github.com/nicrodriguezval/grpc/models"
	"github.com/nicrodriguezval/grpc/protos/questionpb"
	"github.com/nicrodriguezval/grpc/protos/testpb"
	"github.com/nicrodriguezval/grpc/repository"
	"io"
	"log"
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

func (s *TestServer) TakeTest(stream testpb.TestService_TakeTestServer) error {
	questions, err := s.repo.GetQuestionsPerTest(context.Background(), "t1")
	if err != nil {
		return err
	}

	for _, question := range questions {
		question := &questionpb.Question{
			Id:       question.Id,
			Question: question.Question,
		}

		if err := stream.Send(question); err != nil {
			return err
		}

		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		log.Println("Answer:", msg.GetAnswer())
	}

	return nil
}
