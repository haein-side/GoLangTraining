package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y)) // int -> float64
	var z uint = uint(f)                          // float64 -> unit
	fmt.Println(x, y, z)
}

/*
C와 달리 Go는 다른 type의 요소들 간의 할당에는 명시적인 변환을 필요로 함!
즉, 작은 데이터 타입에서 범위가 큰 데이터 타입으로 데이터 타입 변환 시 자동 현변환이 안되는 것이다.

예를 들어 정수형 int에서 uint로 변환할 때, 암묵적(implicit) 변환이 일어나지 않으므로 uint(i) 처럼 반드시 변환을 지정해 주어야 한다.
만약 명시적 지정이 없이 변환이 일어나면 런타임 에러가 발생한다.

cf.
c 언어에선 암시적인 데이터 타입 변환이 일어난다.

int a = 5;
double b = a;  // 암시적인 타입 변환: int를 double로 변환하여 대입


int a = 5;
float b = 2.5;
float result = a + b;  // 암시적인 타입 변환: int를 float로 변환하여 연산 수행
*/