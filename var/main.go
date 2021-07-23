package main

// 기본적인 출력과 입력에 대한 기능을 제공하는 포멧
import "fmt"

func main(){
	a := 4 // 정수 int
	var b float = 4.5 // float64

	var c int = (int)b //에러 : float64 변수를 int에 대입 불가 
	d := float64(a * b) // 에러 : 다른 타입의 int 변수와 float64 연산 불가 
	var e int64 = 16 
	f := (int64)a * e // 에러 : 같은 정수형이지만 int 와 int64로 타입이 달라서 연산 불가 
}



/*
func main(){
	a := 4 // 정수 int
	var b float = 4.5 // float64

	var c int = b //에러 : float64 변수를 int에 대입 불가 
	d := a * b // 에러 : 다른 타입의 int 변수와 float64 연산 불가 
	var e int64 = 16 
	f := a * e // 에러 : 같은 정수형이지만 int 와 int64로 타입이 달라서 연산 불가 
}


	var a int = 3 // 기본 형태
	var b int // 초기값 생략, 초깃값은 타입별 기본값으로 대체
	var c = 4 // 타입 생략, 변수 타입은 우변 값의 타입이 됨
	d := 5 // 선언 대입문 :=을 사용해서 var 키워드와 타입 생략

	fmt.Println(a,b,c,d)

*/