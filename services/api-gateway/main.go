package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
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

	r := mux.NewRouter()
	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

		conn, err := grpc.Dial("rose-auth:3001", opts...)
		if err != nil {
			log.Panic("Fail to dial")
		}
		defer conn.Close()

		client := pb.NewUsersServiceClient(conn)
		users, err := client.GetUsers(context.Background(), &pb.EmptyRequest{})
		if err != nil {
			log.Panic("Can't get users")
		}
	})

	log.Printf("Service \"api-gateway\", listening on port: %s\n", cfg.Port)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	log.Print("Gracefully stopped")
}
