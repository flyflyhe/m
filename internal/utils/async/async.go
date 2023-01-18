package async

import (
	"net/http"
	"time"
)

type HttpResponse struct {
	Res *http.Response
	Err error
}

func aHttp(url string) chan *HttpResponse {
	httpRes := make(chan *HttpResponse)
	go func() {
		res := &HttpResponse{Res: &http.Response{}, Err: nil}
		client := http.Client{Timeout: 2 * time.Second}

		res.Res, res.Err = client.Get(url)

		httpRes <- res
	}()

	return httpRes
}
