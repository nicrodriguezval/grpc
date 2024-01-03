package main

import (
	"github.com/nicrodriguezval/grpc/database"
	svr "github.com/nicrodriguezval/grpc/server"
	"github.com/nicrodriguezval/grpc/studentpb"
	"github.com/nicrodriguezval/grpc/testpb"
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

	studentServer := svr.NewStudentServer(repo)
	testServer := svr.NewTestServer(repo)

	s := grpc.NewServer()

	studentpb.RegisterStudentServiceServer(s, studentServer)
	testpb.RegisterTestServiceServer(s, testServer)

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}
}
