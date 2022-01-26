package main

import (
	"fmt"
	"strconv"
)

func studyDataType() {
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
}
