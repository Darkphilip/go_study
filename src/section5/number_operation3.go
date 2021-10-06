// 데이터 타입: Numeric 연산(3)
package main

import "math"

func main() {
	// 예제1(오버플로우 에러: 범위 초과)
	var n1 uint8 = math.MaxUint8 + 1

	// 예제2(오버플로우 에러: 범위 미만)
	var n5 uint8 = -1
}
