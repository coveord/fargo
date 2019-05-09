package fargo

import (
	"sync"
)

type roundRobin interface {
	Next() string
	Matches([]string) bool
}

type roundrobinImpl struct {
	urls []string
	mu   *sync.Mutex
	next int
}

func newRoundRobin(urls []string) roundRobin {
	return &roundrobinImpl{
		urls: urls,
		mu:   new(sync.Mutex),
	}
}

func (r *roundrobinImpl) Next() string {
	r.mu.Lock()
	sc := r.urls[r.next]
	r.next = (r.next + 1) % len(r.urls)
	r.mu.Unlock()
	return sc
}

func (r *roundrobinImpl) Matches(other []string) bool {

	if (other == nil) != (r.urls == nil) {
		return false
	}

	if len(other) != len(r.urls) {
		return false
	}

	for i := range r.urls {
		if r.urls[i] != other[i] {
			return false
		}
	}

	return true
}
