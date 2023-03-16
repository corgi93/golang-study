# Golang basic


### 함수
* 반환 값이 2개 이상일 수 있음
```
func multifply(w, h int) (int, int){
    return w * 2, h * 2
}

func main(){
    w, h := multiply(10, 20)
    // 20, 40
    fmt.Print(w,h) 
}
```


* strconv.Atoi 함수로 str -> int로 파싱

```
import (
    "fmt"
    "strconv"
)

// "22"로 받으면 22로 변환
strconv.Atoi(string)

```
### 매개변수 전달 방식
- 데이터를 전달할 때의 방식. Go는 "값에 의한 호출" (call by value)를 기본으로 한다.
참조는 해당 데이터를 저장하는 메모리 주소 값을 전달하는 걸로 이해하면 된다.
    ### call by value (값에 의한 호출)
    값에 의한 호출은 함수를 호출 시 매개변수 값을 복사해서 함수 내부로 전달함.
    
    - 함수 내부에서는 전달된 매개변수의 본래 값을 변경할 수 없음.
    
    - 함수 내부에서 본래 값을 변경하려면 &연산자로 변수의 메모리 주소를 전달해야함
    
    ```
    func increase(i int){
        i = i + 1
    }

    func main() {
        i := 10 
        increase(i)
        fmt.Println(i) // 10
    }
    ```
    ### call by reference (침조에 의한 호출)
    참조에 의한 호출로 매개변수 전달하면 함수에서는 전달한 매개변수의 메모리 주소값(참조값)를 매개변수로 받음.
    
    *연산자를 사용해 매개변수 타입을 포인터로 지정해야함.

    참조 타입인 슬라이스, 맵은 메모리의 참조 값을 전달하는게 기본이다.

    ```
    func increase2(i *int) {
	    *i = *i + 1
    }

    func main(){
        // reference by value
	    rbv := 28
	    increase2(&rbv)
	    fmt.Println("reference by value:: ", rbv)
    }
    ```