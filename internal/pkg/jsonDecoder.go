package pkg

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func GetPostData(body io.Reader, v interface{}, w http.ResponseWriter) (err error) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		http.Error(w, "Could not read body", http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(b, v); err != nil {
		http.Error(w, "Could not process request data", http.StatusBadRequest)
		return
	}

	return
}
