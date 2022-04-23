package main

import (
	"errors"
	"log"
	"sync"
	"time"
)

var (
	ErrTooManyRequests = errors.New("too many requests")
)

type LeakyBucket struct {
	ticker   *time.Ticker
	m        sync.Mutex
	cur, cap int64
}

func NewLeakyBucket(cap int64, d time.Duration) *LeakyBucket {
	lb := &LeakyBucket{
		ticker: time.NewTicker(d),
		m:      sync.Mutex{},
		cap:    cap,
	}

	go func() {
		for _ = range lb.ticker.C {
			lb.reset()
		}
	}()

	return lb
}

func (lb *LeakyBucket) reset() {
	lb.m.Lock()
	defer lb.m.Unlock()
	lb.cur = 0
}

func (lb *LeakyBucket) Incr() (int64, error) {
	lb.m.Lock()
	defer lb.m.Unlock()
	if lb.cap <= lb.cur {
		return 0, ErrTooManyRequests
	}
	lb.cur++
	return lb.cap - lb.cur, nil
}

func (lb *LeakyBucket) Close() {
	lb.ticker.Stop()
}

func main() {
	c := NewLeakyBucket(3, time.Second*5)
	defer c.Close()

	ticker := time.NewTicker(time.Millisecond * 900)
	for _ = range ticker.C {
		d, err := c.Incr()
		log.Printf("%d %v", d, err)
	}
}
