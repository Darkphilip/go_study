// if문1

package main

import "fmt"

func main() {
	// 제어문(조건문)
	// if문: 반드시 boolean 검사 -> 1, 0 (사용 불가: 자동 형 변환 불가)
	// 소괄호 사용 x

	var a int = 20
	b := 20

	// 예제1
	if a >= 15 {
		fmt.Println("15이상")
	}

	// 예제2
	if b >= 25 {
		fmt.Println("25이상")
	}

	// 에러 발생1
	/*
	  if b >= 25
	  {
	    fmt.Println("25이상")
	  }
	*/
	// 에러 발생2
	/*
	  if b >= 25
	    fmt.Println("25이상")
	*/
	// 에러 발생3
	/*
		if c := 1; c {
			fmt.Println("True")
		}
	*/

	// 예제3
	if c := 40; c >= 35 {
		fmt.Println("35이상")
	}

	// c += 20 에러 발생

}
