package async

import (
	"context"
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

func aHttpCtx(ctx context.Context, url string) chan *HttpResponse {
	httpRes := make(chan *HttpResponse)
	go func() {
		res := &HttpResponse{Res: &http.Response{}, Err: nil}
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			res.Err = err
			return
		}

		client := http.Client{Timeout: 2 * time.Second}
		
		res.Res, res.Err = client.Do(request)

		httpRes <- res
	}()

	return httpRes
}
