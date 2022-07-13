package model

import (
	"github.com/pankajyadav2741/ott/pkg/error"
	"net/http"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
	case "DELETE":
	default:
		err := error.HandleError(401, "Unsupported Request Method")
		w.WriteHeader(err.StatusCode)
		w.Write([]byte(err.Message))
	}
}

func HandleContent(w http.ResponseWriter, r *http.Request) {

}
