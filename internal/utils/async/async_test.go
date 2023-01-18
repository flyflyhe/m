package async

import (
	"fmt"
	"log"
	"testing"
)

func Test_aHttp(t *testing.T) {
	resChan := aHttp("http://www.baidu.com")
	defer close(resChan)

	fmt.Println("feature mode")
	res := <-resChan
	if res.Err != nil {
		log.Println(res.Err)
	} else {
		fmt.Println(res.Res.Status)
	}
}
