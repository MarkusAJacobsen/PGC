package internal

import (
	"github.com/MarkusAJacobsen/PGC/internal/pkg"
	pgl "github.com/MarkusAJacobsen/pgl/pkg"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func RequestBuilder(method string, url string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, url, body)
	return
}

func SendRequest(req *http.Request) (res *http.Response, err error) {
	client := http.Client{}
	res, err = client.Do(req)
	return
}

func WriteServerError(w http.ResponseWriter, err error) {
	logrus.Errorln(err)
	w.WriteHeader(500)
	w.Write([]byte("An error occurred"))
	pkg.ReportError(pgl.ErrorReport{Msg: "An error occurred", Err: err.Error()})
}
