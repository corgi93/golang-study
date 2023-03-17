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



        
