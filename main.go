package main

import (
	"github.com/nicrodriguezval/grpc/database"
	"github.com/nicrodriguezval/grpc/protos/questionpb"
	"github.com/nicrodriguezval/grpc/protos/studentpb"
	"github.com/nicrodriguezval/grpc/protos/testpb"
	"github.com/nicrodriguezval/grpc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	list, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := list.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	repo, err := database.NewPostgresRepository("postgres://postgres:admin@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	studentServer := server.NewStudentServer(repo)
	testServer := server.NewTestServer(repo)
	questionServer := server.NewQuestionServer(repo)

	s := grpc.NewServer()

	studentpb.RegisterStudentServiceServer(s, studentServer)
	testpb.RegisterTestServiceServer(s, testServer)
	questionpb.RegisterQuestionServiceServer(s, questionServer)

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}
}
