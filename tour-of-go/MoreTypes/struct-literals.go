package main

import "fmt"

type vertex2 struct {
	X, Y int
}

var (
	v1 = vertex2{1, 2}  // has type Vertex
	v2 = vertex2{X: 1}  // Y:0 is implicit - 인자 명시 가능
	v3 = vertex2{}      // X:0 and Y:0 - default 값으로 초기화 가능
	p  = &vertex2{1, 2} // has type *Vertex - 객체 Vertex를 생성하고 메모리 주소 반환 하여 type *Vertex의 포인터 변수
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
