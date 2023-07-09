package main

import "fmt"

func main() {
	i, j := 42, 2701

	// var p *int = &i
	// 자료형 앞에 *을 붙이면 포인터형 변수
	// &i는 i의 메모리 주소를 나타냄
	// 따라서 p는 i의 메모리 주소를 가리키는 포인터 변수가 됨
	// p를 이용하여 i의 값을 가져오거나 변경 가능
	// *p : 역참조 (포인터가 가리키는 메모리 위치의 값 가져옴)
	// *p = 10 은 p가 가리키는 메모리 위치에 값을 저장 가능

	p := &i // point to i

	// *포인터_변수명 : 포인터형 변수에서 값 가져오기
	fmt.Println(*p) // read i through the pointer

	// p가 가리키는 메모리 위치에 21을 저장 가능
	*p = 21        // set i through the pointer
	fmt.Println(i) // see the new value of i

	// p = &j 를 하면 p가 가리키는 메모리 위치가 2701임
	// 따라서 *p는 2701임
	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}