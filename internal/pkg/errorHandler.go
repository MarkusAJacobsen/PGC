package pkg

import (
	"bytes"
	"encoding/json"
	pgl "github.com/MarkusAJacobsen/pgl/pkg"
	"github.com/sirupsen/logrus"
	"net/http"
)

const PGLAddress = "localhost:3333"

func ReportError(errRep pgl.ErrorReport) {
	b, err := json.Marshal(errRep)
	if err != nil {
		logrus.Errorln(err.Error())
		return
	}

	r := bytes.NewReader(b)
	http.NewRequest("POST", PGLAddress, r)
}
