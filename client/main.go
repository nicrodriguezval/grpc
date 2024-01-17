package main

import (
	"github.com/nicrodriguezval/grpc/protos/enrollmentpb"
	"github.com/nicrodriguezval/grpc/protos/questionpb"
	"github.com/nicrodriguezval/grpc/protos/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:5060", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer func() {
		if err := cc.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	testClient := testpb.NewTestServiceClient(cc)
	getTest(testClient)
	takeTest(testClient)

	questionClient := questionpb.NewQuestionServiceClient(cc)
	createQuestions(questionClient)

	enrollmentClient := enrollmentpb.NewEnrollmentServiceClient(cc)
	getStudentsPerTest(enrollmentClient)
}
