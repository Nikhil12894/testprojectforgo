package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Nikhil12894/testprojectforgo/pkg/reading"
)

func main() {

	fmt.Println("server started on port 8080...")

	// router := rest.InitHandler()

	router1 := reading.InitHandler()

	log.Fatal(http.ListenAndServe(":8080", router1))
}
