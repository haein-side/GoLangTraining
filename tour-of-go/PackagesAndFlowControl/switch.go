package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	/*
		runtime.GOOS는 runtime 패키지에 정의되어 있는 상수이고
		운영체제를 확인할 때 사용한다. 운영체제에 따라
		darwin, freebsd, linux, windows를 반환하며 이 값으로 분기 처리를 한다.
	*/
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows
		fmt.Printf("%s. \n", os)
	}
}

/*
Go 언어의 Switch문 특징
- switch문 뒤에 expression 없을 수도 있음
- case 뒤에 expression 가능
- no default fall through
- type switch
*/
