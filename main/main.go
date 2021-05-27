package main

import (
	"log"

	"github.com/Nikhil12894/testprojectforgo/pkg/config"
	"github.com/Nikhil12894/testprojectforgo/pkg/http/rest"
	"github.com/Nikhil12894/testprojectforgo/pkg/redis"
	"github.com/valyala/fasthttp"
)

func main() {
	configuration, err := config.FromFile("main/config.json")
	if err != nil {
		log.Fatal(err)
	}

	service, err := redis.New(configuration.Redis.Host, configuration.Redis.Port, configuration.Redis.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer service.Close()

	router := rest.New(configuration.Options.Schema, configuration.Options.Prefix, service)

	log.Fatal(fasthttp.ListenAndServe(":"+configuration.Server.Port, router.Handler))
}
