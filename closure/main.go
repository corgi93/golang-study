package main

import (
	"fmt"
	"strings"
)

// 확장자를 만들어주는 함수
// return을 익명함수로 받음(클로저)
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		// name에서 suffix포함 되있는지 validation체크
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// f함수를 매개변수로 전달
func callback(y int, f func(int, int)) {
	f(y, 2) // add(10,2)을 호출
}

func add(a, b int) {
	fmt.Printf("%d + %d = %d \n", a, b, a+b) // 1 + 2 = 3
}

func main() {
	fmt.Println("----")
	addZip := makeSuffix(".zip")
	addTgz := makeSuffix(".tar.gz")
	fmt.Println(addTgz("go1.5.1.src"))
	fmt.Println(addZip("go1.5.1.windows-amd64"))

	fmt.Println("----")
	callback(10, add)
}
