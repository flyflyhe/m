package tokenBucket

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateTokenBucket(t *testing.T) {
	tb := CreateTokenBucket(10, 2)
	tb.Start()

	ctxTop, cancelTop := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelTop()

	for i := 0; i < 1000; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		token := tb.GetToken(ctx)
		t.Log(token)
		assert.NotEmpty(t, token)
		cancel()
	}

	for {
		select {
		case <-ctxTop.Done():
			return
		default:
		}
	}

}
