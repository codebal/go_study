package basic

import "fmt"

func studyCondition() {
	if k := 0; k == 1 {
		println("k==", k)
	} else if k > 1 {
		println("k>1")
	} else {
		println("what is k??", k)
	}

	var name string
	category := 1
	switch x := category << 2; x * 2 {
	case 1:
		name = "test1"
	case 2:
		name = "test2"
	case 3, 4:
		name = "test3, 4"
	default:
		name = "default"
	}
	println("name=", name)

	grade := 90
	print("grade is ")
	switch {
	case grade >= 90:
		println("A")
	case grade >= 80:
		println("B")
	default:
		println("No Hope")
	}

	var v interface{} = 1i
	fmt.Printf("v=%v\n", v)
	switch v.(type) {
	case int:
		println("type is int")
	case string:
		println("type is string")
	case bool:
		println("type is bool")
	default:
		println("unknown", v)
	}

	//fallthrough
	val := 10
	switch {
	case val >= 10:
		println("val > 10")
		fallthrough
	case val >= 9:
		println("val > 9")
		fallthrough
	case val >= 5:
		println("val > 5")
		fallthrough
	default:
		println("end of check")

	}

}
