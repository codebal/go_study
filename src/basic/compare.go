package basic

import "fmt"

func StudyCompare() {

	sp1 := "가나다라"
	sp2 := "abcd"
	sp3 := "ABCD"
	//sp4 := "ABCDE"

	fmt.Println(sp1 >= sp2)
	fmt.Println(sp2 >= sp3)
	fmt.Println(sp3 == "ABCD")

	f1 := 0.99991
	f2 := 0.99992
	f3 := 0.9999999999999999999991
	f4 := 0.9999999999999999999992

	fmt.Println("float 오차범위", f1 == f2)
	fmt.Println("float 오차범위", f3 == f4)
}
