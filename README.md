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


-----

## defer

defer 키워드는 함수가 종료되기 전까지 특정 구문의 실행을 지연시켰다가, 함수가 종료되기 직전에 지연시켰던 구문을 수행시킴. 

Java나 C#의 final같은 개념이다. 주로 리소스를 해제시키거나 클렌징 작업이 필요할 때 사용함.

```
package main

import "fmt"

func main() {

	fmt.Println("f1 - start")
	defer f2()
	fmt.Println("f1 - end")

    /*
    실행결과
    f1 - start
    f1 - end
    f2 - deffered
    */
}

func f2() {
	fmt.Printf("f2 - deferred")
}


```
defer을 사용하면 defer로 지정한 각 구문은 stack에 쌓여 있다가 가장 나중에 쌓인 defer구문부터 수행함.

### 사용예시
```
// 파일 stream닫기
file, _ := os.Open(path)
defer file.Close()

// 리소스 잠금 해제하기
mu.Lock()
defer mu.Unlock()

// 데이터베이스 커넥션 닫기
conn, _ = Connection()
defer conn.Close()
```

final보다 defer로 작성해 코드 가독성을 높히고, 특정 리소스의 사용을 해제하는 코드를
defer 구문으로 작성함. 이때는 사용할 리소스를 초기화하는 코드와 해제하는 코드가 함께 있는게 읽기 편하다.

### 내장함수

따로 import하지 않고 사용가능한 내장함수

- close : 고루틴에서 채널 닫을 시 사용
- len : 문자열, 배열, 슬라이스, 맵 채널의 요소 개수 확인 
- cap : 문자열 , 배열, 슬라이스, 맵 채널의 최대 용량 확인
- new : 구조체를 위한 메모리 생성 시 사용
- make : 참조 타입(슬라이스, 맵, 채널)을 위한 메모리 생성할 때 사용
- copy : 배열 및 슬라이스 복사 
- append : 슬라이스 요소 추가
- panic , recover : 에러처리 활용. defer과도 많이 사용
- complex , real, imag : 복소수 처리에 사용


### 클로저(closure)

종종 함수 이름을 정의하지 않고 익명함수로 사용할 때가 있음. Go에서 함수는
일급 객체(first-class object)이므로 변수의 값으로 사용할 수 있음. 다음과 같이 함수를 변수에 할당해 변수처럼 사용이 가능하다.

- 함수를 변수에 할당 가능함
- callback함수를 작성 가능

```
plus := func(x , y int) int {
    return x + y
}
plus(3,5) // 8
```

- 변수에 할당하지 않고 다음과 같이 바로 호출할 수도 있다.
```
func(x, y int) int {
    return x + y
}(3,4)
```



### 패키지 (package)
다른 언어의 모듈, 라이브러리와 유사개념으로 코드를 구조화하고 재사용하기 위한 단위
모든 Go프로그램은 패키지로 구성되고 한 패키지에서 다른 패키지를 import해서 사용.


- 패키지 이름과 디렉토리 이름은 같아야 함.
- 같은 패키지에 있는 소스 파일은 모두 같은 디렉토리에 있어야 함.
- convention
    - 일반적으로 패키지 이름은 '소문자'
    - 소스 파일 하나로 구성된 패키지는 패키지 이름과 소스 파일 이름을 같게 함
- 패키지 종류
    - __실행 가능한 프로그램__
        - 명령 프롬프트에서 명령을내려 실행
        - 패키지 이름이 __main__ 이면 Go가 실행 가능한 프로그램으로 인식!
        - main패키지 빌드하면 디렉토리 이름과 같은 이름으로 실행 파일이 생성되고 실행하면 main패키지의  main()함수를 찾아서 실행함.
        ```

        //  $GOPATH/src/sample/ch2
        $ mkdir pkg

        $ go build

        // 생성된 pkg 실행파일 실행
        $ ./pkg
        Hello World!
        ```
    
    - __라이브러리__
        - 다른 프로그램에서 호출해 사용하도록 연관된 작업을 하는 코드 묶음
        - main패키지 외에는 모두 라이브러리
        - 커스텀 패키지 안에서 $GOPATH/src 디렉토리 기준으로 한 경로로 import해야한다. 
        - 커스텀 패키지의 내부 요소도 __패키지명.식별자__ 로 접근해야함

        src/.../package/lib/lib.go
        ```
        package lib

        func IsDigit(c int32) bool {
            return '0' <= c && c <= '9'
        }
        ```



        src/.../package/pkg/main.go
        ```
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
        ```

- __init()__ 함수
    - 패키기 로드시 가장 먼저 호출되는 함수로, 패키지 초기화 로직이 필요시 개발자의 선택적인 사용
    - react에서 useEffect()같은 ? 
    - Go 프로그램은 항상 main()함수로 시작. main패키지가 다른 패키지 import하고 있다면
    또 import된 패키지에서 또 다른 패키지를 import하고 있다면 import된 패키지를 모두 불러온 후 main()함수 실행한다.

    - 사용법
    ```
    packcage main

    import (
        "fmt"
        "github.com/corgi93/studyGo/package/lib"
    )

    // golang의 rune타입
    var v rune

    func init(){
        v = '1'
    }

    func main(){
        // init()에 v에 '1'할당 후 main()함수 실행
        fmt.Println(lib.IsDigit(v))
    }

    ```
        
<br>
<br>

# Golang 데이터 타입

## 숫자
- 기본적으로 정수는 10진수로 인식
- 8진수는 숫자 0을 붙히고 , 16진수는 숫자 앞에 0x를 붙힌다.
```
a := 365 // 10진수
b := 0555 // 8진수
c := 0x16D // 16진수
```

- 문자 표기
    - Go는 정수 타입과 문자 타입을 구분하지 않음.
    - byte 또는 rune(룬타입)ㅇ으로 문자의 코드 값을 저장해 문자를 표기

    ```
    var ch1 byte = 65
    var ch2 byte = 0101
    var ch3 byte = 0x41

    var ch4 rune = 44032 // 44032 (10진수)
    var ch5 rune = 0126000 // 44032 (8진수)
    var ch6 rune = 0xAC00 // 44032 (16진수)

    fmt.Println("%c %c %c \n", ch1 , ch2, ch3) // A A A
    fmt.Println("%c %c %c \n", ch4, ch5, ch6) // 가 가 가
    ```

- 숫자 연산
산술연산, 비교 연산은 같은 타입끼리 가능. 다른 숫자끼리 연산하려면 타입 변환을 꼭 해줘야 함.

```
i := 100000
j := int16(10000)
k := uint8(100)

// int()로 변환
fmt.Println(i + int(j)) // 110000
fmt.Println(i + int(k)) // 100100
```

## 문자열
문자열(string)은 큰따옴표(")나 백쿼트(`)로 생성. 보통은 큰따옴표로 생성
백쿼트로 생성하면 이스케이프 문자와 줄 바꿈을 무시함.

백쿼트는 HTML, XML같은 문자열 여러줄로 작성하거나 정규식 작성에 좋음

- 문자열에 자주 사용되는 함수
    - len() : 문자열 s의 바이트 수
    - len([]rune(s)) : 문자열 s의 문자 수
    - utf8.RuneCountInString(s) : 문자열 s의 문자 수 
    - strconv.Atoi(s) : 문자열 s를 정수(int)로 변환
    - strconv.Itoa(i) : 정수 i를 문자열로 변환
    

## 문자열과 문자
엄밀히 말하면 Go는 문자(character)타입이 없다. 문자를 표현하려면 정수 타입인 rune(int32의 별칭)으로 문자의 코드값을 사용해야 함

문자는 작은따옴표(')로 감싸 표현함. <br>

문자 'A'의 ASCII코드 값은 65.
문자 'A'의 16진수는 41
문자 'A'의 8진수 101

```
var ch1 = 'A'
var ch2 = 65
var ch3 = '\x41' // 16진수로 표현할 땐 앞에 \x를 붙힘
var ch4 = '\101' // 8진수로 표현할 땐 앞에 \를 붙힘
```


- 문자열 내부 접근
[]로 내부 접근

```
s := "hello"
fmt.Println(s[0]) // h
```

- 문자열 내부 순차적 접근
for ...range 구문을 활용

```
s1 := "hello"
s2 := "안녕하세요"

for i,c := range s1 {
    fmt.Println("%s(%d) \t", c , i)
}

// h(0) e(1) l(2) l(3) o(4)
 
for i,c := range s2 {
    fmt.Println("%s(%d) \t ", c, i)
}
// 안(0) 녕(3) 하(6) 세(9) 요(12)

```

- 인덱스로 문자열 접근

인덱스로 문자열 내부를 접근할 땐 []rune타입으로 변환 후 접근하는 게 안전함
```
s1 := "hello"
s2 := "안녕하세요"
r1 := []rune(s1)
r2 := []rune(s2)

// s1: h e l l o
fmt.Println("s1: %c %c %c %c %c",r1[0],r1[1],r1[2],r1[3],r1[4])

// s2: 안 녕 하 세 요
fmt.Println("s2: %c %c %c %c %c",r2[0],r2[1],r2[2],r2[3],r2[4])
```


- 문자열 변환
```
[]rune(string) // string을 유니코드 문자의 코드값 배열로 변환
[]byte(string) // string을 바이트 배열로 변환
string(char) // 유니코드 문자의 코드값 배열을 문자열로 변환
string(i) // 유니코드의 코드 포인트 i를 무낮열로 변환 (i가 65라면 "A"로 변환)
```

- 문자열 추출

문자열의 부분을 추출
```
s:="hello"
fmt.Println(s[1:2]) //e    :1번째부터 2-1번째 바이트까지 추출 
fmt.Println(s[1:]) // ello  :1번째 바이트부터 마지막 바이트까지 추출
fmt.Println(s[:2]) // he   :2-1번째 까지 추출 

```

## 배열 & 슬라이스

배열 슬라이스 비교

|배열|슬라이스|
|---|---|
|고정 길이값 | 가변적인 길이|
|값 타입 (value type) | 참조 타입 (reference type)
|값에 의한 호출(call by value)로<br>값 전체를 복사해서 전달해야함 | 참조에 의한 호출(call by reference)로 <br> 참조(주소) 값만 전달함
|요소의 타입이 비교 연산자로 비교할 수 있는 타일입 때 <br> 배열 전체에 대해서도 비교연산자로 비교가능 | 비교 연산자 사용불가!


배열보단! 슬라이스가 기능이 풍부해 사용하기도 쉽다! 실제 개발 시 반드시 배열을 사용하는 특별한 상황이 아니라면 슬라이스를 사용하자!


- 배열

초깃값 없으면 제로값으로 초기화함
```
[길이]타입
[길이]타입{초깃값}
[...]타입{초깃값} // ...으로 설정시 지정된 요소의 개수로 배열의 길이를 할당


var arr [5]int // 길이 5인 int형 배열
b := [3]int{1,2,3} // 선언과 동시에 값 초기화
c := [3]int{1,2}
d := [...]int{4,5,6,7,8} // ...을 사용해 배열의 길이 지정

```

- 슬라이스

make함수로 슬라이스 생성시 주어진 길이만큼 제로값으로 초기화된 배열을 내부 메모리에 생성하고 그것의 참조를 반환함. __append()__ 로 슬라이스의 확장할 수 있는 최대 길이가 용량(capacity)이다<br>
append()로 슬라이스 확장 시 슬리이스 길이가 용량을 넘어서면 내부 용량이 증가된 새 배열을 생성함.
```
[]타입
[]타입{초깃값}
make([]타입, 길이 , 용량)
make([]타입, 길이)

var a []int      // int형 슬라이스 선언. 길이,용량은 0으로 지정
var b:= []int{}  // int형 슬라이스 선언. 길이,용량은 0으로 지정
var c:= []int{1,2,3}  // 슬라이스 선언과 동시에 초기화
var d:= [][]int{      // 다차원 슬라이스
    {1,2,},
    {4,5,6}
}
e := make([]int , 0) // make함수로 길이, 용량이 0인 슬라이스 생성
f := make([]int , 3, 5) // make함수로 길이 3, 용량이 5인 슬라이스 생성
```

- for ...range 루프
배열/슬라이스에 for ...range 루프를 사용해 전체 요소를 반복하여 각 요소의 intdex와 값을 얻어올 수 있음

```
numbers := []int{3, 4, 5, 6}
for idx, value := range numbers {
    fmt.Println("index:", idx, "value:", value)
}

/*
index: 0 value: 3
index: 1 value: 4
index: 2 value: 5
index: 3 value: 6
*/

```

- 부분 슬라이스 추출
```
s := []int{1,2,3,4,5,6,7}
fmt.Println(s , "=>", s[:3] , s[3:5] , s[5:], s[:])

// 출력
// [1 2 3 4 5 6 7] => [1 2 3] [4 5] [6 7] [1 2 3 4 5 6 7]
```

- 슬라이스 변경

슬라이스에 새로운 요소 추가하거나 다른 슬라이스를 추가할 때 append()사용<br>
슬라이스 각 요소를 개별로 추가할 땐 ... 연산자를 사용

```
// append(원본 슬라이스, 추가할 요소)
ns1 := []int{1,2,3}
ns2 := []int{6,7,8}
ns3 := []int{8,9,10,11}

ns1 = append(ns1, 4, 5)
ns1 = append(ns1, ns2...)
ns1 = append(ns1, ns3[1:3]...)

fmt.Println(ns1) // [1 2 3 4 5 6 7 8 9 10]
```

- 슬라이스 삽입

슬라이스 마지막이 아닌 처음이나 중간에 값을 삽입할 경우. 기본 함수로 제공되는 게 없으므로 직접 구현해야함. insert함수를 만들어보자.

```

// (슬라이스 s , 삽입할 슬라이스 int[] , 추가할 인덱스 int)
func insert(s, slice int[], index int) int[] {
    return append(s[:index], append(slice, s[index:]...)...)
}

// append()함수 사용하지 않고 make()로 원본 슬라이스 길이와 삽입할 슬라이스 길이를 합한 
// 길이로 슬라이스 생성 후, copy()함수로 각 위치에 요소들을 직접 복사함
func insert2(s, slice int[], index int) int[] {
    res := make([]int, len(s) + len(slice)) // 기존 슬라이스 + 삽입할 슬라이스 길이 
    position := copy(res , s[:index])
    position += copy(res[position:] , slice)
    copy(res[position:], s[index:])
    return res
}

func main(){
    s := []int{1,2,3,4,5}

    s = insert(s, []int{-3,-2}, 0)
    fmt.Println(s) // s:  [-3 -2 1 2 3 4 5]
    
    s = insert(s, []int{0}, 2)
    fmt.Println(s) // s: [-3 -2 0 1 2 3 4 5]

    s = insert(s, []int{6, 7} , len(s))
    fmt.Println(s) //  s: [-3 -2 0 1 2 3 4 5 6 7]
}
```


## 맵(Map)

Go에선  Map의 키 타입은 