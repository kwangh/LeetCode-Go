package leetcode

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	obj := Constructor()
	fmt.Println(obj.Search("a"))
	obj.Insert("apple")
	param_2 := obj.Search("apple")
	if param_2 == false {
		t.Errorf("obj.Search(\"apple\")=%t want=%t", param_2, !param_2)
	}
	param_3 := obj.Search("app")
	if param_3 == true {
		t.Errorf("obj.Search(\"app\")=%t want=%t", param_3, !param_3)
	}
	param_4 := obj.StartsWith("app")
	if param_4 == false {
		t.Errorf("obj.StartsWith(\"app\")=%t want=%t", param_4, !param_4)
	}
	obj.Insert("app")
	param_5 := obj.Search("app")
	if param_5 == false {
		t.Errorf("obj.Search(\"app\")=%t want=%t", param_5, !param_5)
	}
}
