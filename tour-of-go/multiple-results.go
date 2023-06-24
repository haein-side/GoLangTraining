package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x // 두 개의 string 반환
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

/*
한 함수는 몇 개의 결과든 반환 가능
return문을 사용하려면 반드시 Return type을 함수 선언부에 명시해야 함
(생략 불가)
*/
