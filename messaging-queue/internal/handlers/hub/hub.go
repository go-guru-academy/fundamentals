package hub

import (
	"errors"
	"sync"
	"time"
)

type Hub struct {
	messages  chan []byte
	rateCount int
	lock      sync.Mutex
}

const (
	RATE_LIMIT    = 10000
	MESSAGE_LIMIT = 1000000
)

func Init() *Hub {
	h := &Hub{
		messages: make(chan []byte, MESSAGE_LIMIT),
	}
	go h.limit()
	return h
}

func (h *Hub) limit() {
	for {
		h.lock.Lock()
		h.rateCount = 0
		h.lock.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func (h *Hub) Enqueue(message []byte) error {
	select {
	case h.messages <- message:
		h.lock.Lock()
		defer h.lock.Unlock()
		if h.rateCount == RATE_LIMIT {
			return errors.New("error: rate limit exceeded")
		}
		h.rateCount++
	default:
		return errors.New("error: maximum messages in queue")
	}
	return nil
}

func (h *Hub) Dequeue() ([]byte, error) {
	select {
	case b := <-h.messages:
		h.lock.Lock()
		h.rateCount--
		h.lock.Unlock()
		return b, nil
	default:
		return nil, errors.New("error: no messages in queue")
	}
}
