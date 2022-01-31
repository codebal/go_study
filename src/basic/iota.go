package basic

import "fmt"

func StudyIota() {
	const (
		JAN = iota + 1
		FEB
		MAR
		APR
		MAY
		JUN
		JUL
		AUG
		SET
		OCT
		NOV
		DEC
	)
	fmt.Println("JAN=", JAN)
	fmt.Println("APR=", APR)
	fmt.Println("DEC=", DEC)

	const (
		c0 = iota * 10
		c1
		c2
		c3
	)
	fmt.Println(c0, c1, c2, c3)

}
