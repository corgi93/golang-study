package main

import (
	"fmt"

	"github.com/corgi93/studyGo/package/lib"
)

func main() {
	fmt.Println("test...")
	fmt.Println(lib.IsDigit('1'))
	fmt.Println(lib.IsDigit('a'))
}
