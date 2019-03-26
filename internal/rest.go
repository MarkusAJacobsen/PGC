package internal

import (
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
