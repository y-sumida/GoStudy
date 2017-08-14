package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	countWords := map[string]int{}

	for _, words := range strings.Fields(s) {
		countWords[words] += 1
	}

	return  countWords 
}

func main() {
	wc.Test(WordCount)
}
