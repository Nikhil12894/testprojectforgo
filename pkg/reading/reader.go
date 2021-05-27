package reading

import "github.com/gorilla/mux"

func InitHandler() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/reader/", welcomeHandler()).Methods("GET")

	return router
}
