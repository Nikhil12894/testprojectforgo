package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Nikhil12894/testprojectforgo/pkg/http/rest"
)

func main() {

	fmt.Println("server started on port 8080...")

	router := rest.InitHandler()

	log.Fatal(http.ListenAndServe(":8080", router))
}
