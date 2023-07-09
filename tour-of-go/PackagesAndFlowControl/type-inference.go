package main

import "fmt"

func main() {
	v := 0.23
	fmt.Printf("v is of type %T\n", v) // float64 : 부동 소수점 숫자 리터럴은 기본적으로 float64 타입으로 추론됨 / float32(v) 해야 float32 타입 데이터가 됨!
}

/*
타입 유추
:= 혹은 var = 표현을 이용해 명시적인 type을 정의하지 않고 변수를 선언할 때,
그 변수 type은 오른편에 있는 값으로부터 유추됩니다.


하지만, 오른 편에 type이 정해지지 않은 숫자 상수가 올 때,
새 변수는 그 상수의 정확도에 따라 int 혹은 float64, complex128 이 됩니다.
*/
