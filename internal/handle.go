package internal

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func SetUpRouter() (r *mux.Router) {
	r = mux.NewRouter()

	r.HandleFunc("/", baseHandle)
	r.HandleFunc("/plant", plantHandle)

	return
}

func baseHandle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}
