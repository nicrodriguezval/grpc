package main

import (
	"context"
	"github.com/nicrodriguezval/grpc/protos/enrollmentpb"
	"github.com/nicrodriguezval/grpc/protos/questionpb"
	"github.com/nicrodriguezval/grpc/protos/testpb"
	"io"
	"log"
	"sync"
	"time"
)

// unary
func getTest(client testpb.TestServiceClient) {
	req := &testpb.GetTestRequest{
		Id: "t1",
	}

	res, err := client.GetTest(context.Background(), req)
	if err != nil {
		log.Fatalln("error while calling GetTest RPC", err)
	}

	log.Println("Response from GetTest:", res)
}

// client streaming
func createQuestions(client questionpb.QuestionServiceClient) {
	questions := []*questionpb.Question{
		{
			Id:       "q1",
			TestId:   "t1",
			Question: "What is the capital of Spain?",
			Answer:   "Madrid",
		},
		{
			Id:       "q1",
			TestId:   "t1",
			Question: "What is the capital of France?",
			Answer:   "Paris",
		},
		{
			Id:       "q3",
			TestId:   "t1",
			Question: "What is the capital of Germany?",
			Answer:   "Berlin",
		},
	}

	stream, err := client.CreateQuestions(context.Background())
	if err != nil {
		log.Fatalln("error while calling CreateQuestions RPC:", err)
	}

	for _, question := range questions {
		err := stream.Send(question)
		if err != nil {
			log.Fatalln("error while sending question:", err)
		}

		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("error while receiving response from CreateQuestions:", err)
	}

	log.Println("Response from CreateQuestions:", res)
}

// server streaming
func getStudentsPerTest(client enrollmentpb.EnrollmentServiceClient) {
	req := &enrollmentpb.GetStudentsPerTestRequest{
		TestId: "t1",
	}

	stream, err := client.GetStudentsPerTest(context.Background(), req)
	if err != nil {
		log.Fatalln("error while calling GetStudentsPerTest RPC:", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln("error while receiving response from GetStudentsPerTest:", err)
		}

		log.Println("Response from GetStudentsPerTest:", res)
	}
}

// bidirectional streaming
func takeTest(client testpb.TestServiceClient) {
	// the testId is set in the remote procedure
	reqs := []*testpb.TakeTestRequest{
		{
			Answer: "Madrid",
		},
		{
			Answer: "Paris",
		},
		{
			Answer: "Berlin",
		},
	}

	stream, err := client.TakeTest(context.Background())
	if err != nil {
		log.Fatalln("error while calling TakeTest RPC:", err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for _, req := range reqs {
			err := stream.Send(req)
			if err != nil {
				log.Fatalln("error while sending request to TakeTest:", err)
			}

			time.Sleep(time.Second)
		}

		err := stream.CloseSend()
		if err != nil {
			log.Fatalln("error while closing stream:", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatalln("error while receiving response from TakeTest:", err)
			}

			log.Println("Response from TakeTest:", res)
		}
	}()

	wg.Wait()
}
