package internal

import (
	"io"
	"net/http"
	"pgc/internal/pkg"
)

func plantHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		plant := pkg.Plant{}
		pkg.GetPostData(r.Body, &plant, w)

	}
	io.WriteString(w, "Plant endpoint reached")
}

func addPlant(plant pkg.Plant) {}
