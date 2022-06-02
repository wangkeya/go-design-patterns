package pubsub

import (
	"fmt"
	"sync"
	"time"
)

type publisher struct {
	timeout     time.Duration
	subscribers []subscriber
}

type topicFilterFunc func(v interface{}) bool

type subscriber struct {
	buffer int
	ch     chan interface{}
	tf     topicFilterFunc
}

func newPublisher(timeout time.Duration) *publisher {
	return &publisher{timeout: timeout}
}

func newSubscriber(buf int, tf topicFilterFunc) *subscriber {
	return &subscriber{buffer: buf, ch: make(chan interface{}, buf), tf: tf}
}

func (s *subscriber) receive() {
	for msg := range s.ch {
		fmt.Println(msg)
	}
}

func (p *publisher) addSubscribe(sub subscriber) {
	p.subscribers = append(p.subscribers, sub)
}

func (p *publisher) close() {
	for i, _ := range p.subscribers {
		close(p.subscribers[i].ch)
		p.subscribers = append(p.subscribers[0:i], p.subscribers[i+1:]...)
	}
}
func (p *publisher) publish(v interface{}) {
	var wg sync.WaitGroup
	for _, sub := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, v, &wg)
	}
}
func (p *publisher) sendTopic(sub subscriber, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if sub.tf == nil || !sub.tf(v) {
		return
	}

	select {
	case sub.ch <- v:
	case <-time.After(p.timeout):
	}
}
