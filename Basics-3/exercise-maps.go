package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	v := make(map[string]int)

	for _, word:= range strings.Fields(s){
		v[word] += 1
	}

	return v
}

func main() {
	wc.Test(WordCount)
}
