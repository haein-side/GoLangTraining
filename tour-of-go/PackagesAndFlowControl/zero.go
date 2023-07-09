package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s) // 0 0 false ""
}

/*
형식 지시자(format specifier)

%v: %v는 값(value)을 기본 형식으로 출력합니다. 이 형식 지시자는 변수의 타입에 따라 적절한 문자열로 변환하여 출력합니다. 예를 들어, 정수, 부동 소수점, 문자열, 배열, 맵 등의 값이 주어지면 해당 값을 기본 형식으로 출력합니다. 이는 Go의 기본 형식 출력 방식입니다.

%q: %q는 문자열 값을 따옴표로 둘러싸여 있는 형태로 출력합니다. 이 형식 지시자는 문자열 값을 출력할 때 사용됩니다. 문자열 안에 있는 따옴표도 표시하기 위해 이스케이프(Escape)를 적용합니다. 예를 들어, %q를 사용하여 문자열 "Hello, World!"를 출력하면 "Hello, World!"와 같이 따옴표로 둘러싸여 있는 형태로 출력됩니다.

*/
