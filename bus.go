package eventbus

import (
	"errors"
	"sync"
)

type EventBus struct {
	subNode map[string]*node
	rw      sync.RWMutex
}

func NewEventBus() EventBus {
	return EventBus{
		subNode: make(map[string]*node),
		rw:      sync.RWMutex{},
	}
}

func (b *EventBus) Subscribe(topic string, sub Sub) {
	b.rw.Lock()
	if n, ok := b.subNode[topic]; ok {
		b.rw.Unlock()
		n.rw.Lock()
		defer n.rw.Unlock()
		n.subs = append(n.subs, sub)
		return
	}
	defer b.rw.Unlock()
	n := NewNode()
	b.subNode[topic] = &n
	n.subs = append(n.subs, sub)
}

func (b *EventBus) UnSubscribe(topic string, sub Sub) {
	b.rw.Lock()
	if n, ok := b.subNode[topic]; ok && n.SubsLen() > 0 {
		b.rw.Unlock()
		b.subNode[topic].RemoveSub(sub)
		return
	}
	b.rw.Unlock()
}

func (b *EventBus) Publish(topic string, msg interface{}) error {
	b.rw.Lock()
	if n, ok := b.subNode[topic]; ok {
		b.rw.Unlock()
		n.rw.Lock()
		defer n.rw.Unlock()
		go func(subs []Sub, msg interface{}) {
			for _, sub := range subs {
				sub.receive(msg)
			}
		}(n.subs, msg)
		return nil
	}
	defer b.rw.Unlock()
	return errors.New("topic not exists")
}

func (b *EventBus) PubFunc(topic string) func(msg interface{}) {
	return func(msg interface{}) {
		b.Publish(topic, msg)
	}
}

func (b *EventBus) SubsLen(topic string) (int, error) {
	b.rw.Lock()
	if n, ok := b.subNode[topic]; ok {
		b.rw.Unlock()
		n.rw.RLock()
		defer n.rw.RUnlock()
		return n.SubsLen(), nil
	}
	defer b.rw.Unlock()
	return 0, errors.New("topic not exists")
}
