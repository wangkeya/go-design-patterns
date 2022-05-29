package observer

import (
	"fmt"
	"testing"
)

func TestObserver(t *testing.T) {
	var sub = Subject{}
	for i := 0; i < 2; i++ {
		obs := &alarmPlugin{
			id:   i,
			name: "alarm",
			msg:  "hello, alarm",
		}
		sub.attach(obs)
	}
	for i := 2; i < 4; i++ {
		obs := &metricPlugin{
			id:   i,
			name: "metric",
			msg:  "hello, metric",
		}
		sub.attach(obs)
	}
	sub.printObserver()
	sub.notify("start test")
	sub.detach(1)
	sub.printObserver()
	fmt.Println("end all")

}
