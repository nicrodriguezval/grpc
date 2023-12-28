package main

import (
	"github.com/nicrodriguezval/grpc/database"
	svr "github.com/nicrodriguezval/grpc/server"
	"github.com/nicrodriguezval/grpc/studentpb"
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

	server := svr.NewServer(repo)

	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, server)
	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}
}
