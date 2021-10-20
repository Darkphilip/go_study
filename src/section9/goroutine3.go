// 고루틴(Goroutine) 기초(3)
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func exer(name int) {
	r := rand.Intn(100)
	fmt.Println(name, "start:", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, ">>>>>", r, i)
	}
	fmt.Println(name, "end:", time.Now())

}

func main() {
	// 고루틴(Goroutine)
	// 멀티코어(다중CPU) 최대한 활용

	runtime.GOMAXPROCS(runtime.NumCPU())                      // 현 시스템의 CPU 코어 개수 반환 후 설정
	fmt.Println("Current System CPU:", runtime.GOMAXPROCS(0)) //설정값 출력

	// 예제1
	fmt.Println("Main Routine Start:", time.Now())
	for i := 0; i < 100; i++ {
		go exer(i) // 고루틴 100개 생성
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Main Routine End:", time.Now())

}
