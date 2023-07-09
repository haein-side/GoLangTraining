package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	q                 = "해인"
)

func main() {
	q = "종인"
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Println(q)
}

/*
int 와 uintptr type은 보통 32-bit 시스템에서는 32 bit, 64-bit 시스템에서는 64 bit의 길이입니다.
정수 값이 필요할 때에는 특정한 이유로 크기를 정해야하거나 unsigned 정수 type을 사용해야하는 게 아니라면 int 를 사용해야합니다.

명시적 타입간 변환만 가능
Go에서 타입간 변환은 명시적으로 지정해 주어야 한다는 점인데,
예를 들어 정수형 int에서 uint로 변환할 때, 암묵적(implicit) 변환이 일어나지 않으므로 uint(i) 처럼 반드시 변환을 지정해 주어야 한다. 만약 명시적 지정이 없이 변환이 일어나면 런타임 에러가 발생한다.
*/