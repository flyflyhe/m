package tokenBucket

import (
	"context"
	"log"
	"sync"
	"time"
)

type Token struct {
	Val  string
	Time int64
}

type TokenBucket struct {
	queue     chan *Token
	size      int
	timer     *time.Ticker
	duration  int64
	onceStart sync.Once
}

func CreateTokenBucket(size, duration int) *TokenBucket {
	return &TokenBucket{queue: make(chan *Token, size), size: size, timer: time.NewTicker(time.Duration(duration) * time.Second)}
}

func (this *TokenBucket) Start() {
	this.onceStart.Do(func() {
		go func() {
			for {
				select {
				case <-this.timer.C:
					log.Println("tick")
					for i := 0; i < this.size-len(this.queue); i++ {
						this.queue <- &Token{Val: "", Time: time.Now().Unix()}
					}
				default:
				}
			}
		}()
	})
}

func (this *TokenBucket) GetToken(context context.Context) *Token {
	for {
		select {
		case <-context.Done():
			return nil
		case t := <-this.queue:
			return t
		default:
		}
	}
}
