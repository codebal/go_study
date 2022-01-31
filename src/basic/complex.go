package basic

import "fmt"

func studyComplex_1() {
	var c1 complex64 = 1 + 2i
	var c2 complex64 = 3i
	var c3 complex64 = 2.0

	fmt.Println(c1 + c2)
	fmt.Println(c1 - c3)

	var f1 = real(c1)
	var f2 = imag(c1)
	fmt.Println(f1)
	fmt.Println(f2)

	var c4 = complex(f1, f2)
	fmt.Println(c4)
}

func studyComplex_2() {
	var c1 complex64 = 1 + 1i
	var c2 complex64 = 2 + 3i

	fmt.Println("c1 * c2 = ", c1*c2)
}

func StudyComplex() {
	studyComplex_2()
}
