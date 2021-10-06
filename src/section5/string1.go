// 데이터 타입: 문자열(1)

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 문자열
	// 큰 따옴표"", 백스쿼트 ``
	// Golang: 문자 char 타입 존재하지 않음 -> rune(int32)로 문자 코드 값으로
	// 문자: '' 로 작성
	// 자주 사용하는 escape: \\, \', \", \a(콘솔벨), \b(백스페이스), \f(쪽바꿈), \r(복귀), \t(탭)..

	// 예제1
	var str1 string = "c:\\go_study\\src\\" // -> c:\go_study\src\
	str2 := `c:\go_study\src\`

	fmt.Println("ex1:", str1)
	fmt.Println("ex1:", str2)

	// 예제2
	var str3 string = "Hello, World!"
	var str4 string = "안녕하세요"
	var str5 string = "\ud55c\uae00"

	fmt.Println()
	fmt.Println("ex2:", str3)
	fmt.Println("ex2:", str4)
	fmt.Println("ex2:", str5)

	// 예제3
	// 길이(바이트 수)
	fmt.Println("ex3:", len(str3))
	fmt.Println("ex3:", len(str4))

	// 예제4
	// 길이(실제 길이)
	fmt.Println("ex4:", utf8.RuneCountInString(str3))
	fmt.Println("ex4:", utf8.RuneCountInString(str4))
	fmt.Println("ex4:", len([]rune(str4))) // len으로 실제 길이 구하는 법
}
