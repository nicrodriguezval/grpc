package server

import (
	"context"
	"github.com/nicrodriguezval/grpc/models"
	"github.com/nicrodriguezval/grpc/protos/enrollmentpb"
	"github.com/nicrodriguezval/grpc/protos/studentpb"
	"github.com/nicrodriguezval/grpc/repository"
	"io"
	"time"
)

type EnrollmentServer struct {
	repo repository.Repository
	enrollmentpb.UnimplementedEnrollmentServiceServer
}

func NewEnrollmentServer(repo repository.Repository) *EnrollmentServer {
	return &EnrollmentServer{repo: repo}
}

func (s *EnrollmentServer) Enroll(stream enrollmentpb.EnrollmentService_EnrollServer) error {
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&enrollmentpb.EnrollmentResponse{
				Success: true,
			})
		}
		if err != nil {
			return err
		}

		enrollment := &models.Enrollment{
			Id:        msg.GetId(),
			StudentId: msg.GetStudentId(),
			TestId:    msg.GetTestId(),
		}

		err = s.repo.Enroll(context.Background(), enrollment)
		if err != nil {
			return stream.SendAndClose(&enrollmentpb.EnrollmentResponse{
				Success: false,
			})
		}
	}
}

func (s *EnrollmentServer) GetStudentPerTest(
	req *enrollmentpb.GetStudentPerTestRequest,
	stream enrollmentpb.EnrollmentService_GetStudentPerTestServer,
) error {
	students, err := s.repo.GetStudentsPerTest(context.Background(), req.GetTestId())
	if err != nil {
		return err
	}

	for _, student := range students {
		student := &studentpb.Student{
			Id:   student.Id,
			Name: student.Name,
			Age:  student.Age,
		}

		time.Sleep(time.Second)
		err := stream.Send(student)
		if err != nil {
			return err
		}
	}

	return nil
}
