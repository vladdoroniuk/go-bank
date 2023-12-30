package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v10"
	"github.com/gorilla/mux"
	"github.com/vladdoroniuk/rose/services/api-gateway/handlers"

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

	r := mux.NewRouter().StrictSlash(true)
	api := r.PathPrefix("/api").Subrouter()

	handlers.AuthHandlers(api)

	go func() {
		srv := http.Server{
			Addr:    ":" + cfg.Port,
			Handler: r,
		}
		if err := srv.ListenAndServe(); err != nil {
			log.Panic("Can't start server")
		}
	}()

	log.Printf("Service \"api-gateway\", listening on port: %s\n", cfg.Port)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	log.Print("Gracefully stopped")
}
