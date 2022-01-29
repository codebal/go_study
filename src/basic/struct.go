package basic

import "fmt"

type person struct {
	name string
	age  int
}

type dict struct {
	data map[int]string
}

func newDict() *dict {
	d := new(dict)
	d.data = make(map[int]string)
	return d
}

func StudyStruct() {
	{
		p := person{}

		p.name = "Lee"
		p.age = 10

		fmt.Printf("%v\n", p)
	}
	{
		dic := newDict()
		dic.data[1] = "A"
		fmt.Printf("dic = %v\n", dic)
	}
}
