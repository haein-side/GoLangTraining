package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // len=0 cap=6 []

	// Extend its length.
	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	/* 용량(capacity)을 초과하는 경우 현재 용량의 2배에 해당하는
	새로운 Underlying array (주: 아래 내부구조 참조) 을 생성하고
	기존 배열 값들을 모두 새 배열에 복제한 후 다시 슬라이스를 할당 */

	s = append(s, 2, 3, 4, 5, 6)
	printSlice(s) // len=9 cap=12 [2 3 5 7 2 3 4 5 6]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
