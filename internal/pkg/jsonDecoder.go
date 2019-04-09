package pkg

import (
	"encoding/json"
	pgl "github.com/MarkusAJacobsen/pgl/pkg"
	"io"
	"io/ioutil"
	"net/http"
)

const CouldNotRead = "Could not read body"
const CouldNotProcess = "Could not process request data"

func GetPostData(body io.Reader, v interface{}, w http.ResponseWriter) (err error) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		http.Error(w, CouldNotRead, http.StatusBadRequest)
		ReportError(pgl.ErrorReport{Msg: CouldNotRead, Err: err.Error()})
		return
	}

	if err = json.Unmarshal(b, v); err != nil {
		http.Error(w, CouldNotProcess, http.StatusBadRequest)
		ReportError(pgl.ErrorReport{Msg: CouldNotProcess, Err: err.Error()})
		return
	}

	return
}
