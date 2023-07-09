package main

import "fmt"

var i, j int = 1, 2
var k = 3 // type 추론 가능
// q := 2 // 함수 바깥에서는 := 사용 불가

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, k, python, java)
}
