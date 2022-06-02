package pubsub

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"
	"time"
)

func Test_PubSub(t *testing.T) {
	p := newPublisher(100 * time.Millisecond)
	defer p.close()
	s := newSubscriber(10, func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})
	p.addSubscribe(*s)
	p.publish("hello, golang")
	p.publish("hello, php")
	p.publish("hello, java")

	go func() {
		s.receive()
	}()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sig:
		fmt.Println("finished!!!")

	}

}
