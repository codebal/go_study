package basic

import (
	"fmt"
	"strconv"
)

func studyDataType_1() {
	var (
		b  bool
		in int
		s  string
		f  float32
		c  complex64 //허수
		by byte
		r  rune
	)

	println(b, in, s, f, c, by, r)

	str1 :=
		`
		this is raw string literal by quote
		line break without \n
		`
	println(str1)

	str2 := fmt.Sprintf(`
		with format
		number = %d
	`, 100)
	println(str2)

	//change type
	var i int = 100
	var u uint = uint(i)
	var fl float32 = float32(i)
	println(u, fl)

	str3 := "ABC"
	bytes := []byte(str3)
	str4 := string(bytes)
	println(bytes, str4)

	convertInt, err := strconv.ParseInt("12345", 10, 32)
	if err == nil {
		println("convertInt= ", convertInt+1)
	}

	var r1 rune = '한'
	r2 := '\ud55c'
	fmt.Printf("rune for unicode like hangle = %v, %v \n", r1, r2)
}

func studyDataType_2() {
	fmt.Println("8 / 3 = ", 8/3)
	fmt.Println("5 / 9 = ", 5/9)
	fmt.Println("9 % 5 = ", 9%5)
}

func StudyDataType() {
	studyDataType_2()
}
