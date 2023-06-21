package main

import "fmt"

// 패키지 전역 변수
// main 패키지 안에서는 다 쓸 수 있음
var g int = 10

func main() {
	var m int = 20

	{
		var s int = 50
		fmt.Println(m, s, g)
	}

	// 변수는 속해있는 중괄호에만 살아있음
	// m = s + 20 // s undeclared - 변수의 범위를 벗어났기 때문
}
