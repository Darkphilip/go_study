package main

import "fmt"

type Person struct {
	name   string
	age    int
	weight float32
}

func main() {
	p := Person{name: "윤병찬", age: 24, weight: 88.8}

	fmt.Println(p)
	fmt.Printf("Person 구조체형 변수p가 저장된 메모리: %p\n", &p)
	fmt.Printf("p.name이 저장된 메모리 주소: %p\n", &p.name)
}
