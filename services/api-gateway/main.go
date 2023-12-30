package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/vladdoroniuk/rose/proto_gen/auth"
)

type config struct {
	Port string `env:"PORT"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Panic("Can't parse env")
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(":"+cfg.Port, opts...)
	if err != nil {
		log.Panic("Fail to dial")
	}
	defer conn.Close()

	client := pb.NewUsersServiceClient(conn)
	users, err := client.GetUsers(context.Background(), &pb.EmptyRequest{})
	if err != nil {
		log.Panic("Can't get users")
	}

	fmt.Println(users)

	log.Printf("Service \"api-gateway\", listening on port: %s\n", cfg.Port)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	log.Print("Gracefully stopped")
}
