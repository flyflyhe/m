package lru

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

type Data struct {
	prev *Data
	next *Data
	key  string
	val  string
}

type Lru struct {
	start bool
	head  *Data
	cap   int
	m     sync.Map
	dirty sync.Map
	q     chan *Data
	stop  chan struct{}
	num   int
}

func (lru *Lru) show() {
	root := lru.head.next
	for root != nil {
		fmt.Println(root.val)
		root = root.next
	}
}

func CreateLru(cap, qSize int) *Lru {
	return &Lru{
		head:  &Data{},
		cap:   cap,
		m:     sync.Map{},
		dirty: sync.Map{},
		q:     make(chan *Data, qSize),
		stop:  make(chan struct{}),
	}
}

func (lru *Lru) Start() {
	lru.start = true
	go func() {
		for {
			select {
			case data := <-lru.q:
				if target, ok := lru.m.Load(data.key); !ok {
					data.prev = lru.head
					data.next = lru.head.next
					lru.num++
					if lru.num > lru.cap { //容量已经达到 删除末尾
						root := lru.head
						prev := new(Data)
						for root.next != nil {
							prev = root
							root = root.next
						}
						prev.next = nil
						lru.m.Delete(root.key)
						lru.num--
					}

					if lru.head.next != nil {
						lru.head.next.prev = data
					}

					lru.m.Store(data.key, data)
					lru.dirty.Delete(data.key)
					lru.head.next = data
				} else {
					log.Println("get", data.key)
					oldData := target.(*Data)
					if oldData.next != nil {
						oldData.next.prev = data.prev
					}
					oldData.prev.next = data.next

					data.next = lru.head.next
					data.prev = lru.head
					lru.head.next = data
				}

			case <-lru.stop:
				lru.start = false
				break
			default:
			}
		}
	}()
}

func (lru *Lru) Stop() {
	lru.stop <- struct{}{}
}

func (lru *Lru) Put(key, val string) {
	if !lru.start {
		panic("必须先启动")
	}
	lru.dirty.Store(key, val)

	lru.q <- &Data{key: key, val: val}
}

func (lru *Lru) Get(key string) (string, error) {
	if !lru.start {
		panic("必须先启动")
	}
	if v, ok := lru.dirty.Load(key); ok {
		return v.(string), nil
	}

	if v, ok := lru.m.Load(key); ok {
		data := v.(*Data)
		lru.q <- data //命中放到头部
		return data.val, nil
	}

	return "", errors.New("not found")
}
