package main

import (
	"CalculatorService/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	proto.UnimplementedCalculatorServiceServer
}


func (s *Server) PrimeNumberDecomposition(req *proto.IntRequest,stream proto.CalculatorService_PrimeNumberDecompositionServer)  error{
	number := int(req.Number)
	for number%2 == 0 {
		res := &proto.IntResponse{Number: 2, OptionalError: ""}
		if err := stream.Send(res); err != nil {
			log.Fatalf("error while sending stream responses: %v", err.Error())
		}
		number = number / 2
	}


	for i := 3; i*i <= number; i = i + 2 {
		// while i divides n, append i and divide n
		for number%i == 0 {
			res := &proto.IntResponse{Number: int64(i), OptionalError: ""}
			if err := stream.Send(res); err != nil {
				log.Fatalf("error while sending stream responses: %v", err.Error())
			}
			number = number / i
		}
	}

	if number > 2 {
		res := &proto.IntResponse{Number: int64(number), OptionalError: ""}
		if err := stream.Send(res); err != nil {
			log.Fatalf("error while sending stream responses: %v", err.Error())
		}
	}

	return nil
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:59751")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(s, &Server{})
	log.Println("Server is running on port:59751")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
