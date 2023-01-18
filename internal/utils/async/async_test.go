package async

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func Test_aHttp(t *testing.T) {
	resChan := aHttp("https://www.baidu.com")
	defer close(resChan)

	fmt.Println("feature mode")
	res := <-resChan
	if res.Err != nil {
		log.Println(res.Err)
	} else {
		fmt.Println(res.Res.Status)
	}
}

func Test_aHttpCtx(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	resChan := aHttpCtx(ctx, "https://www.google.com")
	defer close(resChan)

	fmt.Println("feature mode")
	time.Sleep(time.Second)

	cancel()

	res := <-resChan
	if res.Err != nil {
		log.Println(res.Err)
	} else {
		fmt.Println(res.Res.Status)
	}
}
