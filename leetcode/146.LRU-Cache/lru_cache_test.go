package leetcode

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	param_1 := obj.Get(1)
	if param_1 != 1 {
		t.Errorf("want: 1 instead got %d", param_1)
	}
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
}
