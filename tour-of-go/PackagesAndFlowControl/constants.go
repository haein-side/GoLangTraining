package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	const ( // var() 처럼 한 번에 여러 개의 변수 선언 가능!
		visa   = "visa"
		master = "masterCard"
		amx    = "american express"
	)

	const (
		Apple = iota
		Grape
		Orange
	)

	// visa = "hey"
	// 상수이므로 다시 초기화 불가능
	// cannot assign to visa (untyped string constant "visa")

	fmt.Println(visa)
}

/*
상수는 := 를 사용하지 못한다.
그 이유는 상수의 경우 선언과 초기화가 동시에 한 줄에서 이루어져야 한다.
값의 불변성을 유지하기 위해서이다.

따라서 상수는 = 만 사용하여 선언과 동시에 초기화가 일어나야 한다.
이렇게 선언된 상수는 값을 변경할 수 없고, 컴파일 타임에 초기화되어 고정된 값을 가지게 됩니다.

상수는 변수처럼 선언되지만 const 키워드와 함께 선언됩니다.
상수는 character 혹은 string, boolean, 숫자 값이 될 수 있습니다.
Go 언어의 상수는 컴파일 타임에 평가되는 값 -> 프로그램 실행 중에 동적으로 변경될 수 없으며, 컴파일 시에 값을 결정합니다.
*/

/*
Go 언어는 타입 안정성을 가진 언어!
변수의 타입은 명시적으로 선언되거나 컴파일러가 타입을 추론하여 결정.

타입 안정성
프로그래밍 언어에서 변수나 표현식의 타입이 명확하게 지정되고, 타입 규칙을 준수해야 하는 특성을 말합니다.
타입 안정성은 잠재적인 타입 오류를 방지하고 프로그램의 안정성과 신뢰성을 높이는데 도움을 줍니다.
(타입 안정성 있는 언어에서는 변수 타입을 미리 지정함!)

타입 안정성이 있는 언어에서는 변수의 타입을 선언하거나 할당 시에 명시적으로 지정하며,
해당 타입과 일치하지 않는 타입의 값을 할당하거나 사용하려고 하면 "컴파일 오류"가 발생합니다. 이는 변수와 값을 사용하는 시점에서 타입 호환성을 검사하여 타입 오류로 인한 문제를 "사전에 방지합"니다.
*/