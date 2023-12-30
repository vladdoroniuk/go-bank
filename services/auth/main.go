package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v10"
	"github.com/gorilla/mux"
)

type config struct {
	Port string `env:"PORT"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Panic("Can't parse env vars")
	}

	r := mux.NewRouter().StrictSlash(true)
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test"))
	})

	go func() {
		server := &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: r,
		}
		if err := server.ListenAndServe(); err != nil {
			log.Panic("Can't start server")
		}
	}()

	log.Printf("Service \"auth\", listening on port: %s\n", cfg.Port)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	log.Print("Gracefully stopped")
}
