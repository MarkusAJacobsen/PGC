package pkg

import (
	"bytes"
	"encoding/json"
	pgl "github.com/MarkusAJacobsen/pgl/pkg"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
)

const PGLAddress = "http://127.0.0.1:3333/error"

func ReportError(errRep pgl.ErrorReport) {
	b, err := json.Marshal(errRep)
	if err != nil {
		logrus.Errorln(err.Error())
		return
	}

	u, err := url.Parse(PGLAddress)
	if err != nil {
		panic(err)
		return
	}

	r := bytes.NewReader(b)
	resp, err := http.Post(u.String(), "application/json", r)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	logrus.Info(string(b))
}
