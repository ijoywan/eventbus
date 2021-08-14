## What

a simple event bus in golang

## Use

```golang
package main

import (
	"fmt"
	"github.com/ijoywan/eventbus"
)

// 创建一个事件总线
var bus eventbus.EventBus

func main() {
	bus = eventbus.NewEventBus()

	// 创建订阅器
	sub := eventbus.NewSub()

	exit := make(chan bool)
	bus.Subscribe("topic1", sub)

	// 订阅器处理函数
	go func() {
		msg := sub.Out().(int)
		fmt.Printf("got number:%d\n", msg)
		exit <- true
	}()

	// 发起任务
	bus.Publish("topic1", 7)

	// 等待任务处理结束
	<-exit
}
```

or 

```go
package main

import (
	"fmt"

	"github.com/ijoywan/eventbus"
)

// 创建一个事件总线
var bus eventbus.EventBus

func main() {
	bus = eventbus.NewEventBus()

	// 创建订阅器
	sub := eventbus.NewSub()

	bus.Subscribe("topic1", sub)

	// 发起任务
	bus.Publish("topic1", 7)

	// 等待返回结果
	msg := sub.Out().(int)
	fmt.Printf("got number:%d\n", msg)
}
```

For more details,see [testfile](https://github.com/ijoywan/eventbus/blob/master/bus_test.go)