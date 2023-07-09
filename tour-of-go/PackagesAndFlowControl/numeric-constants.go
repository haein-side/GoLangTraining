package main

import "fmt"

const (
	// 패키지 레벨에서 선언한 상수는 해당 패키지 내부 어디에서든지 사용 가능하다.
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	// 100000 .... 00000 (0이 총 100개) -> 2 ^ 100
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	// 10 -> 2 ^ 1 -> 2
	Small = Big >> 99
)

// 패키지 레벨에서 선언한 변수는 해당 패키지 내부 어디에서든지 사용 가능하다.
var try = 2

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small)) // const Small 인자로 넣어줌
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	const sample = 100

	fmt.Println(sample)
	fmt.Println(try)
}

/*
Go 언어에서는 상수와 변수를 main 함수 내에만 제한하지 않고 어디에서든 선언과 초기화할 수 있습니다.
상수와 변수는 함수 내부뿐만 아니라 패키지 수준에서도 선언할 수 있습니다.

함수와 변수는 함수 내부뿐만 아니라 패키지 수준에서도 선언할 수 있습니다.

패키지 수준에서 선언된 상수와 변수는 해당 패키지 내의 모든 함수에서 사용할 수 있습니다.
이러한 패키지 수준의 상수와 변수는 해당 패키지의 모든 파일에서 공유되며,
패키지 내의 여러 함수에서 값을 공유하고 사용할 수 있습니다.

package main

import "fmt"

const myConstant = 42 // 패키지 수준에서 선언된 상수
var myVariable = "Hello, World!" // 패키지 수준에서 선언된 변수

func main() {
    fmt.Println(myConstant)
    fmt.Println(myVariable)
}
*/
