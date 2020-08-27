package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	fields := strings.Fields(s)
	result := make(map[string]int)

	for _, v := range fields {
		var _, isExist = result[v]
		if isExist {
			result[v] += 1
		} else {
			result[v] = 1
		}
	}

	return result
}

func main() {
	wc.Test(WordCount)
}
