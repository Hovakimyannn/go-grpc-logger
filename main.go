package main

import (
	"context"
	_ "fmt"
	"log"
	"net"

	pb "awesomeProject/logger/gen/logger" // Replace with the path to your generated files

	"google.golang.org/grpc"
)

// Server implements the LoggerServiceServer interface
type Server struct {
	pb.UnimplementedLoggerServiceServer
}

// LogMessage logs a message and returns a response
func (s *Server) LogMessage(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {
	log.Printf("Received message: %s", req.GetMessage())
	return &pb.LogResponse{Status: "Message logged successfully!"}, nil
}

func main() {
	// Create a listener on TCP port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a gRPC server
	grpcServer := grpc.NewServer()

	// Register the LoggerService with the gRPC server
	pb.RegisterLoggerServiceServer(grpcServer, &Server{})

	log.Println("Server is running on port 50051...")
	// Start the server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
