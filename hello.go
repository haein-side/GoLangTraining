package main

// main
// 프로그램 시작 포함하는 패키지
// main () 가지는 패키지, 무조건 1개 (여러 개의 패키지 가질 수 있음)

import "fmt"

// fmt
// 패키지 중 하나를 import해서 가져옴
// import main(x)

func main() {
	// func
	// 함수
	// main
	// 약속어, 프로그램 시작점

	fmt.Printf("hello world")
	// fmt 패키지 안에 포함된 Println이라는 함수 쓰겠다는 것
}

// 프로그램 종료
