package basic

import (
	"fmt"
	"strconv"
)

func studyBit_1() {
	var a int = 6 & 5
	fmt.Println(a)

	fmt.Println(6 & 5)
	fmt.Println(6 | 5)
	fmt.Println(6 ^ 5)
	fmt.Println(6 &^ 5)

	fmt.Println(strconv.ParseInt("10", 2, 32))

	bitOperation("110", "101")
}

func bitOperation(b1, b2 string) {
	i1, _ := strconv.ParseInt(b1, 2, 32)
	i2, _ := strconv.ParseInt(b2, 2, 32)

	fmt.Println(i1, i2)
	fmt.Println(i1, "&", i2, "=", i1&i2)
	fmt.Println(i1, "|", i2, "=", i1|i2)
	fmt.Println(i1, "^", i2, "=", i1^i2)
	fmt.Println(i1, "&^", i2, "=", i1&^i2)
}

func studyBit_2() {
	fmt.Println("3<<4", 3<<4)
	fmt.Println("48>>4", 48>>4)

	fmt.Println("-3<<4", -3<<4)
	fmt.Println("-48>>4", -48>>4)
}

func StudyBit() {
	studyBit_2()
}
