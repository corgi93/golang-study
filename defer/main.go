package main

import "fmt"

func main() {

	fmt.Println("f1 - start")
	defer f2()
	fmt.Println("f1 - end")

	f3()
	fmt.Println()

	b()
}

func f2() {
	fmt.Printf("f2 - deferred")
}

func f3() {
	for i := 0; i < 5; i++ {
		// defer스택에 쌓였다가 맨 위부터 쌓인 부분부터 print
		//  0,1,2,3,4로 쌓이다 4,3,2,1,0 으로 출력
		defer fmt.Printf("%d", i)
	}
}

// defer사용해 trace log출력
func enter(s string) string {
	fmt.Println("entering: ", s)
	return s
}

func leave(s string) {
	fmt.Println("leaving: ", s)
}

func a() {
	defer leave(enter("a"))
	fmt.Println("in a")
}

func b() {
	defer leave(enter("b"))
	fmt.Println("in b")
	a()
}
