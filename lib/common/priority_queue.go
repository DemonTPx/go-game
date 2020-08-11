package common

import (
	"container/heap"
	"errors"
)

type PriorityQueue struct {
	itemHeap *itemHeap
	lookup   map[interface{}]*item
}

func NewPriorityQueue() PriorityQueue {
	return PriorityQueue{
		itemHeap: &itemHeap{},
		lookup:   make(map[interface{}]*item),
	}
}

func (p *PriorityQueue) Len() int {
	return p.itemHeap.Len()
}

func (p *PriorityQueue) Insert(v interface{}, priority float64) {
	_, ok := p.lookup[v]
	if ok {
		return
	}

	newItem := &item{
		value:    v,
		priority: priority,
	}
	heap.Push(p.itemHeap, newItem)
	p.lookup[v] = newItem
}

func (p *PriorityQueue) Pop() (interface{}, error) {
	if len(*p.itemHeap) == 0 {
		return nil, errors.New("empty queue")
	}

	item := heap.Pop(p.itemHeap).(*item)
	delete(p.lookup, item.value)
	return item.value, nil
}

func (p *PriorityQueue) UpdatePriority(x interface{}, newPriority float64) {
	item, ok := p.lookup[x]
	if !ok {
		return
	}

	item.priority = newPriority
	heap.Fix(p.itemHeap, item.index)
}

type itemHeap []*item

type item struct {
	value    interface{}
	priority float64
	index    int
}

func (ih *itemHeap) Len() int {
	return len(*ih)
}

func (ih *itemHeap) Less(i, j int) bool {
	return (*ih)[i].priority < (*ih)[j].priority
}

func (ih *itemHeap) Swap(i, j int) {
	(*ih)[i], (*ih)[j] = (*ih)[j], (*ih)[i]
	(*ih)[i].index = i
	(*ih)[j].index = j
}

func (ih *itemHeap) Push(x interface{}) {
	it := x.(*item)
	it.index = len(*ih)
	*ih = append(*ih, it)
}

func (ih *itemHeap) Pop() interface{} {
	old := *ih
	item := old[len(old)-1]
	*ih = old[0 : len(old)-1]
	return item
}
