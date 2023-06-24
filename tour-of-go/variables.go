package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java) // 선언하고 초기화하지 않으면 default 값 : 0, false, ""

	var j, k int = 2, 3
	fmt.Println(j, k)

	var s = "Hi" // 타입 선언 안해도 할당되는 값 보고 타입 추론해줌
	var z = 1
	fmt.Println(s, z)

	var p = 1
	// p := 1 // 변수 p가 이미 선언되었으므로 여기서 또 선언해줄 수 없음
	fmt.Println(p)

	q := 2
	fmt.Println(q)
}

/*
Go에서 변수와 상수는 함수 밖에서도 사용할 수 있다.
함수 밖에서는 := 사용 못하고 var 사용해야 한다.

만약 선언된 변수가 Go 프로그램 내에서 사용되지 않는다면, 에러를 발생시킨다. 따라서 사용되지 않는 변수는 프로그램에서 삭제한다.

동일한 타입의 변수가 복수 개 있을 경우, 변수들을 나열하고 마지막에 타입을 한번만 지정할 수 있다.

할당되는 값을 보고 그 타입을 추론하는 기능이 자주 사용됨

Short Assignment Statement ( := )
함수(func) 내에서만 사용할 수 있으며, 함수 밖에서는 var를 사용해야 한다.
*/
