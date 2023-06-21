package main

import "fmt"

func main() {
	var a int16 = 3456   // a는 int16 타입 - 2바이트 정수
	var b int8 = int8(a) // int16 -> int8

	fmt.Println(a, b) // 3456, -128
	
	// 타입 변환 시 조심할 것!
	// 큰 그릇 -> 작은 그릇 - 나머지 버려짐
	// 작은 그릇 -> 큰 그릇 - 버려지는 거 없고 다 들어감
	// a int16 = 3456
	// 00001101 10000000
	//    X       남음

	// 실수 -> 정수 시엔 소수점 자릿수들은 날아감
}
