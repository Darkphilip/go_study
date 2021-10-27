// Go 에러 처리 기초(3)
package main

import (
	"errors"
	"fmt"
)

func mainerror3() {
	// 에러 처리
	// errors 패키지의 New 메서드를 이용한 에러 생성
	// Errof보다 많이 사용

	var err1 error = errors.New("Error occured - 1")
	err2 := errors.New("Error occured - 2")

	fmt.Println("error1:", err1)
	fmt.Println("error1:", err1.Error())

	fmt.Println("error2:", err2)
	fmt.Println("error2:", err2.Error())

}
