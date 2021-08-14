## What

a simple event bus in golang

## Use

```golang
	bus := NewEventBus()
	sub := NewSub()

	bus.Subscribe("topic1", sub)

	go func() {
		msg := sub.Out().(int)
		if msg != 7 {
			t.Fatalf("got wrong number:%d", msg)
		}
	}()
	bus.Publish("topic1", 7)
	// or use PubFunc to publish to a certain topic
	pubFunc := bus.PubFunc("topic1")
	pubFunc(7)

```
For more details,see [testfile](https://github.com/ijoywan/eventbus/blob/master/bus_test.go)