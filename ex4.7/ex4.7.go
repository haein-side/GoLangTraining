package main

import "fmt"

func main() {
	var a float32 = 1234.523
	var b float32 = 3456.123
	var c float32 = a * b
	var d float32 = c * 3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	// 0.426666334329 but 4.266663e+06 - float32는 소수부 7자리까지만 살아남음
	fmt.Println(d)
	// 잘리기 때문에 연산의 오차가 생김
	// float32 - 소수부 7자리
	// float64 - 소수부 15자리

}
