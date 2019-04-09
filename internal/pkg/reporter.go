package pkg

import (
	"bytes"
	"encoding/json"
	pgl "github.com/MarkusAJacobsen/pgl/pkg"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const PGLAddress = "http://172.19.0.3:6113/report"
const errorPath = "/error"
const trafficPath = "/traffic"

func ReportError(errRep pgl.ErrorReport) {
	b, err := json.Marshal(errRep)
	if err != nil {
		logrus.Errorln(err.Error())
		return
	}

	u, err := url.Parse(PGLAddress+errorPath)
	if err != nil {
		panic(err)
		return
	}

	sendLog(b, u)
}

func TrafficMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := httputil.DumpRequest(r, true)
		if err != nil {
			logrus.Errorln(err)
			return
		}

		t := pgl.TrafficReport{Msg: string(b)}

		b, err = json.Marshal(t)
		if err != nil {
			logrus.Errorln(err.Error())
			return
		}

		u, err := url.Parse(PGLAddress+trafficPath)
		if err != nil {
			panic(err)
			return
		}
		sendLog(b, u)

		next.ServeHTTP(w, r)
	})
}

func sendLog(b []byte, u *url.URL) {
	r := bytes.NewReader(b)
	resp, err := http.Post(u.String(), "application/json", r)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorln(err)
		return
	}
}