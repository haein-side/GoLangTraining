package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

/*

Go 프로그래밍 언어에서 함수는 리턴값이 없을 수도, 리턴값이 하나 일 수도, 또는 리턴값이 복수 개일 수도 있다. C 언어에서 void 혹은 하나의 값만을 리턴하는 것과 대조적으로 Go 언어는 복수개의 값을 리턴할 수 있다.

Go에서 Named Return Parameter도 있다.
리턴되는 값들을 (함수에 정의된) 리턴 파라미터들에 할당할 수 있는 기능이다.
이는 리턴되는 값들이 여러 개일 때, 코드 가독성을 높이는 장점이 있다.

함수 정의에서 return 변수의 명시적 이름과 타입을 지정하고 이를 리턴 가능
반드시 return 써줘야 함

가독성을 왜 높인다는 것이냐?
단지 함수의 signature만을 읽고도 return parameter들을 알 수 있기 때문

:= 사용 불가
Go 컴파일러에 의해 이미 return parameter들이 초기화되었기 때문에 오류가 발생. 따라서 (=)을 사용하여 명명된 반환 매개변수에 값을 할당 가능.

*/
