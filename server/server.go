package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"../../Screen/screenpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sample(ctx context.Context, req *screenpb.SampleRequest) (*screenpb.SampleResponse, error) {
	fmt.Println("Request came in:", req)

	responseValue := req.GetCalculation().GetFirstInt() + req.GetCalculation().GetSecondInt()
	fmt.Println(responseValue)
	res := &screenpb.SampleResponse{
		Result: responseValue,
	}
	return res, nil
}

func (*server) Speak(ctx context.Context, req *screenpb.SpeakRequest) (*screenpb.SpeakResponse, error) {
	fmt.Println("Speak Request came in:", req)

	responseValue := req.GetSpeakPhrase()
	fmt.Println(responseValue)
	res := &screenpb.SpeakResponse{
		Result: responseValue,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	screenpb.RegisterScreenServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
