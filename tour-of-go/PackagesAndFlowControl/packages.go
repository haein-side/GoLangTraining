package main

import (
	"fmt"
	"math/rand"
)

// 패키지를 추가하기 위해서는 디렉토리 경로의 마지막 이름을 사용하는 게 규칙
// rand가 패키지명

func main() {
	fmt.Println("My favorite number is", rand.Intn(10)) // 0에서 10 이하의 난수 발생
}
