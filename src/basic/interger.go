package basic

import (
	"fmt"
	"unsafe"
)

func studyInteger_1() {
	var i1 int8
	fmt.Println(unsafe.Sizeof(i1)) //sizeof는 바이트 수
	var i2 int16
	fmt.Println(unsafe.Sizeof(i2))
	var i3 int32
	fmt.Println(unsafe.Sizeof(i3))
	var i4 int64
	fmt.Println(unsafe.Sizeof(i4))
}

func studyInteger_2() {
	var i1 uint8 = 0xFF
	_ = i1
	//fmt.Println(i1)

	fmt.Println("int8 최솟값:", -0x80, " 최댓값:", 0x7F)
	fmt.Println("int16 최솟값:", -0x8000, " 최댓값:", 0x7FFF)
	fmt.Println("int32 최댓값:", -0x80000000, " 최댓갑:", 0x7FFFFFFF)
	fmt.Println("int64 최댓값:", -0x8000000000000000, " 최댓값:", 0x7FFFFFFFFFFFFFFF)
}

func StudyInteger() {
	studyInteger_2()
}
