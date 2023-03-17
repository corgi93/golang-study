package main

import "fmt"

func increase(i int) {
	i = i + 1
}

func increase2(i *int) {
	*i = *i + 1
}

type Person struct {
	name string
	age  int
}

func (p Person) toString() {
	p.age++
	print(&p, "\n")
}

func main() {
	// call by value
	cbv := 17
	increase(cbv)
	fmt.Println("call by value:: ", cbv)

	// reference by value
	rbv := 28
	increase2(&rbv)
	fmt.Println("reference by value:: ", rbv)

	print("할당된 두 변수의 주소값 \n")
	p1 := Person{name: "hyeokjin", age: 31}
	p2 := Person{name: "beomjin", age: 32}
	print("p1 주소: ", &p1, "\n", "p2 주소: ", &p2, "\n")

	// 위의 &p1주소와 다름
	p1.toString()
	p2.toString()

	fmt.Println(p1.age)
	fmt.Println(p2.age)
}
