package observer

import (
	"fmt"
	"strconv"
)

type Observer interface {
	update(msg msgType) error
	getId() int
	getMsg() msgType
}

type msgType string

type alarmPlugin struct {
	id   int
	name string
	msg  msgType
}

func (o *alarmPlugin) getId() int {
	return o.id
}

func (o *alarmPlugin) getMsg() msgType {
	return o.msg
}

func (o *alarmPlugin) update(msg msgType) error {
	o.msg = msg
	fmt.Printf("alarm id %d updated\n", o.id)
	return nil
}

type metricPlugin struct {
	id   int
	name string
	msg  msgType
}

func (o *metricPlugin) getId() int {
	return o.id
}

func (o *metricPlugin) getMsg() msgType {
	return o.msg
}

func (o *metricPlugin) update(msg msgType) error {
	o.msg = msg
	fmt.Printf("metric id %d updated\n", o.id)
	return nil
}

type Subject struct {
	observers []Observer
}

func (s *Subject) attach(observer Observer) {
	fmt.Printf(">>> attach id %d observer\n", observer.getId())
	s.observers = append(s.observers, observer)
}

func (s *Subject) detach(obsId int) error {
	fmt.Printf(">>> detach id %d observer\n", obsId)
	var flag = false
	for i, observer := range s.observers {
		if obsId == observer.getId() {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			flag = true
			break
		}
	}
	if flag {
		return nil
	}
	return fmt.Errorf("%d observer not found\n", obsId)
}

func (s *Subject) notify(msg msgType) {
	for _, observer := range s.observers {
		observer.update(msg)
	}
}

func (s *Subject) printObserver() {
	var obsIds []msgType
	for _, observer := range s.observers {
		id := strconv.Itoa(observer.getId())
		obsIds = append(obsIds, msgType(id)+"-\""+observer.getMsg()+"\"")
	}
	fmt.Printf("print observer: %v \n", obsIds)
}
