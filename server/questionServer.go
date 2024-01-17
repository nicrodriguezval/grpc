package server

import (
	"context"
	"github.com/nicrodriguezval/grpc/models"
	"github.com/nicrodriguezval/grpc/protos/questionpb"
	"github.com/nicrodriguezval/grpc/repository"
	"io"
)

type QuestionServer struct {
	repo repository.Repository
	questionpb.UnimplementedQuestionServiceServer
}

func NewQuestionServer(repo repository.Repository) *QuestionServer {
	return &QuestionServer{repo: repo}
}

func (s *QuestionServer) CreateQuestions(stream questionpb.QuestionService_CreateQuestionsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&questionpb.CreateQuestionResponse{
				Success: true,
			})
		} else if err != nil {
			return err
		}

		question := &models.Question{
			Id:       msg.GetId(),
			Answer:   msg.GetAnswer(),
			Question: msg.GetQuestion(),
			TestId:   msg.GetTestId(),
		}

		err = s.repo.CreateQuestion(context.Background(), question)
		if err != nil {
			return stream.SendAndClose(&questionpb.CreateQuestionResponse{
				Success: false,
			})
		}
	}
}
