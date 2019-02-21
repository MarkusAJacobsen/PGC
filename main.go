package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Panic("Port not sat")
	}
	fmt.Println(port)

	r := mux.NewRouter()
	r.HandleFunc("/", handleRoutes)

	if err := http.ListenAndServe("0.0.0.0:"+port, r); err != nil {
		fmt.Println(err)
	}
}

func handleRoutes(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}
