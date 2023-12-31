package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

/*
Go 언어의 defer 키워드
특정 문장 혹은 함수를 나중에 (defer를 호출하는 함수가 리턴하기 직전에) 실행하게 함.
일반적으로 defer는 C#, Java 같은 언어에서의 finally 블럭처럼 마지막에 clean up 작업을 위해 사용됨.
*/
