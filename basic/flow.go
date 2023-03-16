package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 흐름제어 - 조건문 if
	var aa int = 20
	var bb int = 30

	if aa > bb {
		fmt.Println("aa is bigger than bb")
	}
	fmt.Println("bb is bigger then aa")

	// 흐름제어 - switch
	i := -2
	switch i {
	case -1, -2:
		fmt.Println(i, "는 음수 입니다.")
	case 1, 2:
		fmt.Println(i, "는 양수 입니다.")
	}

	// 중간에 끊을라면 fallthrough
	switch i {
	case -2:
		fmt.Println("fallthrough 실행")
		fallthrough
	case -1:
		fmt.Println("fallthrough 후 로직 실행")
	}

	level := 2
	switch level {
	case 1:
		// level이 1일 때 수행
		fmt.Println("do first!")
		fallthrough
	case 2:
		// level이 1또는 2일 때 수행
		fmt.Println("do second!")
		fallthrough
	case 3:
		// level이 1,2 또는 3일때 수행
		fmt.Println("do third")
	}

	// for 초기화구문; 조건식 ;후속작업 {}
	for i := 10; i < 15; i++ {
		fmt.Print(i, " ,")
	}

	fmt.Println("-----")
	// slice나 map을 사용가능 함.
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range numbers {
		fmt.Print(i, ", ")
	}

	/*
		for문에 레이블로 식별자 붙일 수 있음. 콜론(:)으로 끝나는 문자가 있으면 레이블로 인식함
		continue, break, 레이블을 함께 사용해 유연하게 제어가능
	*/

	x := 7
	table := [][]int{{1, 5, 9}, {2, 6, 5, 13}, {5, 3, 7, 4}}
	found := false

	for row := 0; row < len(table); row++ {
		for col := 0; col < len(table[row]); col++ {
			if table[row][col] == x {
				found = true
				fmt.Printf("found %d (row : %d , col : %d)\n", x, row, col)
				break
			}
		}
		if found {
			break
		}
	}

	// refactoring - 레이블 활용, range키워드로 리팩토링
	y := 13
LOOP:
	for row, rowValue := range table {
		fmt.Println("row: ", row)
		fmt.Println("rowValue:", rowValue)
		for col, colValue := range rowValue {
			if colValue == y {
				fmt.Printf("found! %d (row : %d , col : %d)\n", x, row, col)
				break LOOP
			}
		}
	}

	myFunc("check!", 12, 23)

	w, h := multiply(20, 30)
	fmt.Println(w, h)

	// strconv 패키지의 AtoI 함수를 사용해 문자열을 정수로 변환하여 결과와 에러 상태를 반환.
	displayInt("22")

}

func myFunc(s string, integers ...int) {
	// ...으로 가변인자로 매개변수 개수 정해져 있지 않고 유동적으로 변할 때 사용하면 유용함
	fmt.Println(s)

	for _, value := range integers {
		fmt.Println("value::", value)
	}
}

// go에는 값을 두 개 이상 반환하는 함수가 많음. (수행 결과와 에러 상태를 반환)
func multiply(width, height int) (int, int) {
	return width * 2, height * 2
}

func displayInt(s string) {

	if value, err := strconv.Atoi(s); err != nil {
		fmt.Printf("%s는 정수가 아닙니다. \n", s)
	} else {
		fmt.Printf("정수 값은 %d 입니다. \n", value)
	}
}
