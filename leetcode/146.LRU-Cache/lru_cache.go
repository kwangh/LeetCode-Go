package leetcode

import "container/list"

type LRUCache struct {
	cap int
	m   map[int]*list.Element
	l   *list.List
}

type Item struct {
	key, val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		m:   make(map[int]*list.Element, capacity),
		l:   list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.l.MoveToFront(v)
		return v.Value.(*Item).val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.m[key]; ok {
		v.Value.(*Item).val = value
		this.l.MoveToFront(v)
		return
	}

	if this.l.Len() >= this.cap {
		element := this.l.Back()
		this.l.Remove(element)
		delete(this.m, element.Value.(*Item).key)
	}

	element := this.l.PushFront(&Item{key: key, val: value})
	this.m[key] = element
}
