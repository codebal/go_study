package basic

import (
	"fmt"
)

func studyArray() {
	{
		var a [3]int
		a[0] = 1
		a[1] = 2
		a[2] = 3
		fmt.Println(a)
	}

	{
		var a1 = [3]int{1, 2, 3}
		var a2 = [...]int{1, 2, 3, 4}
		var a3 = []int{1, 2, 3, 4, 5, 6}

		fmt.Println("a1", len(a1), cap(a1))
		fmt.Println("a2", len(a2), cap(a2))
		fmt.Println("a3", len(a3), cap(a3))
	}

	{
		//다차원
		var a = [][]int{
			{2, 3, 4},
			{4, 5, 6},
		}
		fmt.Println("a = ", a, len(a[0]), cap(a))
	}
}
