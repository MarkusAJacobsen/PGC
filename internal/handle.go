package internal

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func SetUpRouter() (r *mux.Router) {
	r = mux.NewRouter()

	r.HandleFunc("/", baseHandle)
	pss := r.PathPrefix("/plant").Subrouter()
	pss.HandleFunc("", plantHandle).Methods(http.MethodGet, http.MethodPost)
	pss.HandleFunc("/batch", plantBatchHandle).Methods(http.MethodPost)

	r.HandleFunc("/user", userHandle).Methods(http.MethodPost, http.MethodPut)
	r.HandleFunc("/user/{uIdToken}/projects", projectHandle).Methods(http.MethodGet)
	r.HandleFunc("/user/{uIdToken}/project/{pId}", userProjectHandler).Methods(http.MethodGet)
	r.HandleFunc("/project", projectHandle).Methods(http.MethodPost, http.MethodPut)
	r.HandleFunc("/project/{pId}", projectHandle).Methods(http.MethodDelete)

	gss := r.PathPrefix("/guide").Subrouter()
	gss.HandleFunc("", GuideHandle).Methods(http.MethodPost)

	return
}

func baseHandle(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello world")
}
