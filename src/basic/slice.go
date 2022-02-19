package basic

import "fmt"

func studySlice_1() {
	{
		var a []int
		a = []int{1, 2, 3}
		a[1] = 10
		fmt.Println(a)
	}

	{
		//make
		s := make([]int, 5, 10)
		println(len(s), cap(s), s)
		fmt.Println(len(s), cap(s), s)
		//fmt.Printf("printf = %s, %v", s, s)
	}

	{
		var s []int
		if s == nil {
			println("Nil Slice", s)
		}
		println(len(s), cap(s))
	}

	{
		s := []int{0, 1, 2, 3, 4, 5}
		s = s[2:5]
		fmt.Println(s)
	}

	{
		s := []int{0, 1, 2, 3, 4, 5}
		s = s[2:5]
		fmt.Println(s)
		s = s[1:]
		fmt.Println(s)
	}

	{
		s := []int{0, 1}
		fmt.Println(s)
		s = append(s, 2)
		fmt.Println(s)
		s = append(s, 3, 4, 5)
		fmt.Println(s)
	}

	{
		//slice는 cap를 초과할때 기존의 2배가 됨
		sliceA := make([]int, 0, 3)

		for i := 0; i <= 15; i++ {
			sliceA = append(sliceA, i)
			fmt.Println(len(sliceA), cap(sliceA))
		}
		fmt.Println(sliceA)
	}

	{
		sliceA := []int{1, 2, 3}
		sliceB := []int{4, 5, 6}

		sliceA = append(sliceA, sliceB...)

		var func1 = func(n ...int) int {
			s := 0
			for _, num := range n {
				s += num
			}
			return s
		}
		fmt.Println("func1 = ", func1())

		fmt.Println(sliceA)
	}

	{
		//copy
		source := []int{0, 1, 2}
		println("source = ", len(source), cap(source))
		target := make([]int, len(source), cap(source)*2)
		println("target = ", len(target), cap(target))
		copy(target, source)
		fmt.Println(target)
		println("target = ", len(target), cap(target))
	}

	{
		//slice len cap 구조
		arr1 := []int{0, 1, 2, 3, 4, 5}
		fmt.Println("arr1 = ", arr1, len(arr1), cap(arr1))
		arr2 := arr1[2:4]
		fmt.Println("arr2 = ", arr2, len(arr2), cap(arr2))
		arr3 := append(arr2, 1, 4)
		fmt.Println("arr3 = ", arr3, len(arr3), cap(arr3))
		fmt.Println(arr1, arr2, arr3)
	}
}

func studySlice_2() {
	//arr1 := make([]int, 5)
	arr1 := []int{}
	fmt.Println(arr1, len(arr1), cap(arr1))
	arr2 := append(arr1, 1, 2, 3)
	fmt.Println(arr2, len(arr2), cap(arr2))
}

//slice is pointer
//array is value
func studySlice_3() {
	arr1 := []int{1, 2, 3}
	arr2 := arr1
	arr2[0] = 10
	fmt.Println(arr1, arr2)

	arr3 := [3]int{1, 2, 3}
	arr4 := arr3
	arr4[0] = 10
	fmt.Println(arr3, arr4)
}

//copy는 덮어쓰기
//cap을 늘려주지 않는다
func studySlice_4() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6, 7, 8}
	fmt.Println(arr1, arr2)
	copy(arr1, arr2)
	fmt.Println(arr1, arr2)
}

func studySlice_5() {
	arr1 := [5]int{1, 2, 3, 4, 5} //array or slice
	arr2 := arr1[:3]              //slice
	arr3 := arr2
	arr3[0] = 10

	fmt.Println(arr1, arr2, arr3)
}

func StudySlice() {
	studySlice_5()
}
