package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env/v10"
)

type config struct {
	Port string `env:"PORT"`
}

func main() {
	fmt.Println("Auth service")

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Panic("Can't parse env vars")
	}

	done := make(chan bool)
	go func() {
		if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
			log.Panic("Can't start server")
		}
	}()

	fmt.Printf("Listening on port: %s\n", cfg.Port)
	<-done
}
