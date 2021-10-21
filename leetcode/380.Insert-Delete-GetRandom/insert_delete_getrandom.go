package leetcode

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	s []int
	m map[int]int
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano())
	return RandomizedSet{nil, make(map[int]int)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.m[val] = len(this.s)
	this.s = append(this.s, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	i, ok := this.m[val]
	if !ok {
		return false
	}

	if i < len(this.s)-1 {
		length := len(this.s)
		last := this.s[length-1]
		this.s[i] = last
		this.m[last] = i
	}
	delete(this.m, val)
	this.s = this.s[:len(this.s)-1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.s[rand.Intn(len(this.s))]
}
