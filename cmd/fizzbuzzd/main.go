package main

import (
	"log"
	"net/http"

	"github.com/JoLePheno/Fizz-Buzz/internal/service"
	"github.com/caarlos0/env"
)

var (
	e Environment
)

type Environment struct {
	HttpServerPort string `env:"HTTP_SERVER_PORT" envDefault:"3000"`
}

func main() {
	if err := env.Parse(&e); err != nil {
		log.Fatal("Failed to parse environment")
	}

	s := service.AlgoService{}
	log.Fatal(http.ListenAndServe(":"+e.HttpServerPort, s.Router()))
}
