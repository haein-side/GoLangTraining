package main

import "fmt"

func main() {
	// 정수 <=> 실수 타입 변환 가변
	a := 3 // var a int64 = 3
	var b float64 = 3.5

	var c int = int(b) // 3.5 -> 3
	d := float64(a) * b

	var e int64 = 7
	f := a * int(e)

	fmt.Println(a, b, c, d, e, f)

	/*
		변수란
		값을 저장하는 메모리 공간을 가리키는 이름
		변수를 통해 값을 찾아갈 수 있다

		a = 20
		메모리 공간(왼쪽)
		mov [0x12afbad], 20
		왼쪽 메모리 공간에 20을 복사한다

		c = a
		공간(왼쪽) 값(오른쪽) -> 값이 공간에 복사된다
		c = 20

		타입이란
		공간 크기
		GO는 강타입 언어 (타입을 강하게 검사!)

		nil
		정의되지 않은 메모리 주소, 메모리에 없는 공간
	*/
}
