package eventbus

type Sub struct {
	out chan interface{}
}

func NewSub() Sub {
	return Sub{
		out: make(chan interface{}),
	}
}

func (s *Sub) receive(msg interface{}) {
	s.out <- msg
}

func (s *Sub) Out() (msg interface{}) {
	return (<-s.out)
}
