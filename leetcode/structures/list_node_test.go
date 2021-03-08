package structures

import (
	"testing"
)

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Ints returns int slices of list node
func TestInts(t *testing.T) {
	test := ListNodes{Val: 1, Next: &ListNodes{Val: 1, Next: &ListNodes{Val: 2, Next: &ListNodes{Val: 3, Next: &ListNodes{Val: 4, Next: &ListNodes{Val: 4}}}}}}
	want := []int{1, 1, 2, 3, 4, 4}
	if !equal(Ints(&test), want) {
		t.Errorf("want: %v, instead got %v\n", want, Ints(&test))
	}
}

func TestLists(t *testing.T) {
	test := []int{1, 1, 2, 3}
	want := &ListNodes{Val: 1, Next: &ListNodes{Val: 1, Next: &ListNodes{Val: 2, Next: &ListNodes{Val: 3}}}}
	testList := Lists(test)
	for testList != nil && want != nil {
		if testList.Val != want.Val {
			t.Errorf("want: %d, instead got %d\n", want.Val, testList.Val)
		}
		testList = testList.Next
		want = want.Next
	}
}
