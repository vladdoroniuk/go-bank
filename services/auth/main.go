package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	pb "github.com/vladdoroniuk/rose/proto_gen/auth"
)

/* type config struct {
	Port string `env:"PORT"`
} */

type Server struct {
	pb.UnimplementedUsersServiceServer
}

func New() *Server {
	return &Server{}
}

func (s *Server) GetUsers(ctx context.Context, req *pb.EmptyRequest) (*pb.GetUsersResponse, error) {
	return &pb.GetUsersResponse{
		Name: "Vlad",
		Age:  21,
	}, nil
}

func main() {
	/* cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Panic("Can't parse env vars")
	} */

	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUsersServiceServer(grpcServer, New())
	log.Println("Service \"auth\", listening on port: 3001")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	//log.Printf("Service \"auth\", listening on port: %s\n", cfg.Port)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	log.Print("Gracefully stopped")
}
