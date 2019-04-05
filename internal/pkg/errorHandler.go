package pkg

import (
	"bytes"
	"encoding/json"
	pgl "github.com/MarkusAJacobsen/pgl/pkg"
	"net/http"
)

const PGLAddress = "localhost:6113"

func reportError(errRep pgl.ErrorReport) (err error) {
	b, err := json.Marshal(errRep)
	if err != nil {
		return err
	}

	r := bytes.NewReader(b)
	http.NewRequest("POST", PGLAddress, r)

	return nil
}
