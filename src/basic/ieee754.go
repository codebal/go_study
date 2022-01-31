package basic

import "fmt"

func studyIEEE754_1() {
	var f1 float32 = 0.1
	var f2 float32 = .1
	var f3 float32 = 1e-1
	fmt.Println(f1, f2, f3)
}

func studyIEEE754_2() {
	var f1 = .0

	for i := 0; i < 10; i++ {
		f1 += 0.1
		fmt.Println("f1 = ", f1)
	}

	fmt.Println(f1)
	fmt.Println(f1 == 1.0)
	const errbound = 1e-14
	fmt.Println((f1 - 1.0) <= errbound)
}

func StudyIEEE754() {
	studyIEEE754_2()
}
