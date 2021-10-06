// 자료형: 맵(4)
package main

import "fmt"

func main() {
	// 맵(Map)
	// 맵 조회 할 경우에 주의할 점

	// 예제1
	map1 := map[string]int{ //int:0, string:"",float:0.0
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	value1 := map1["lemon"]
	value2, f := map1["lemon"]
	value3, t := map1["kiwi"]

	fmt.Println("ex1:", value1)
	fmt.Println("ex1:", value2, f)
	fmt.Println("ex1:", value3, t) // 두 번째 리턴 값으로 키 존재 유무 확인

	// 예제2
	if value, ok := map1["kiwi"]; ok {
		fmt.Println("ex2:", value)
	} else {
		fmt.Println("ex2: kiwi is not exist!")
	}

	if value, ok := map1["banana"]; ok {
		fmt.Println("ex2:", value)
	} else {
		fmt.Println("ex2: banana is not exist!")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("ex2: kiwi is not exist!")
	}

}
