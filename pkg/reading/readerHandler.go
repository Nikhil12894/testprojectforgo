package reading

import (
	"encoding/json"
	"net/http"
)

func welcomeHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("welcome!")
	}
}
