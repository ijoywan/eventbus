package eventbus

import "sync"

type node struct {
	subs []Sub
	rw   sync.RWMutex
}

func NewNode() node {
	return node{
		subs: []Sub{},
		rw:   sync.RWMutex{},
	}
}

func (n *node) SubsLen() int {
	return len(n.subs)
}

func (n *node) RemoveSub(s Sub) {
	lenOfSub := len(n.subs)
	n.rw.Lock()
	defer n.rw.Unlock()
	idx := n.findSubIdx(s)
	if idx < 0 {
		return
	}
	copy(n.subs[idx:], n.subs[idx+1:])
	n.subs[lenOfSub-1] = Sub{}
	n.subs = n.subs[:lenOfSub-1]

}

func (n *node) findSubIdx(s Sub) int {
	for idx, sub := range n.subs {
		if sub == s {
			return idx
		}
	}
	return -1
}
