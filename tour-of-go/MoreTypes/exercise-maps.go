package main

import (
	//"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	words := strings.Fields(s) // 문자열의 공백을 기준으로 단어 분리

	for _, word := range words {
		counts[word]++
	}

	return counts
}

func main() {
	//wc.Test(WordCount)
}
