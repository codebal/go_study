package basic

import "fmt"

type Pointer struct {
	var1 string
	var2 int
}

func studyFunc() {
	msg := "Hello"
	say(msg)

	msg2 := 12345
	say2(&msg2)

	pointer := Pointer{"pointer test", 100}
	say3(&pointer)

	total := sum(1, 2, 3, 4, 5)
	println("total = ", total)

	count2, total2 := sum2(1, 2, 3, 4, 5)
	fmt.Printf("count2=%v, total2=%v\n", count2, total2)

	count3, total3 := sum2(1, 2, 3, 4, 5)
	fmt.Printf("count3=%v, total3=%v\n", count3, total3)
}

func say(msg string) {
	println(msg)
}

func say2(msg *int) {
	//println(*msg)
	fmt.Printf("%v, %v, %v, %v, %v\n", msg, *msg, &msg, *&msg, **&msg)
	//*msg = "Changed"
}

func say3(pointer *Pointer) {
	fmt.Printf("%v, %v, %v, %v, %v\n", pointer, *pointer, &pointer, pointer.var1, pointer.var2)
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func sum2(nums ...int) (int, int) {
	s := 0
	count := 0
	for _, n := range nums {
		s += n
		count++
	}
	//count = len(nums)
	return count, s
}

func sum3(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}
