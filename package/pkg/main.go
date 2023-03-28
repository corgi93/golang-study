package main

import (
	"fmt"

	// mylib 이렇게 alias(별칭)을 줄 수
	mylib "github.com/corgi93/studyGo/package/lib"
)

func main() {
	fmt.Println("test...")
	fmt.Println(mylib.IsDigit('1'))
	fmt.Println(mylib.IsDigit('a'))

	// 소문자로 작성하면 빌드에러
	// 소문자로 작성하면 외부 패키지에서는 접근할 수 없다. (내부 패키지에서만)
	// fmt.Println(lib.isSpace('\t'))

	numbers := []int{3, 4, 5, 6}
	for idx, value := range numbers {
		fmt.Println("index:", idx, "value:", value)
	}

	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(s, "=>", s[:3], s[3:5], s[5:], s[:])

	ns1 := []int{1, 2, 3}
	ns2 := []int{6, 7, 8}
	ns3 := []int{8, 9, 10, 11}

	ns1 = append(ns1, 4, 5)
	ns1 = append(ns1, ns2...)
	ns1 = append(ns1, ns3[1:3]...)

	fmt.Println(ns1) // [1 2 3 4 5 6 7 8 9 10]
}
