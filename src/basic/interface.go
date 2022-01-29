package basic

import (
	"fmt"
	"math"
)

type Shape interface {
	area1() float64
	perimeter1() float64
}

type Rect1 struct {
	width, height float64
}

type Circle1 struct {
	radius float64
}

func (r Rect1) area1() float64      { return r.width * r.height }
func (r Rect1) perimeter1() float64 { return 2 * (r.height * r.height) }

func (c Circle1) area1() float64      { return math.Pi * c.radius * c.radius }
func (c Circle1) perimeter1() float64 { return 2 * math.Pi * c.radius }

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area1()
		fmt.Println(a)
	}
}

func Assertion() {
	var a interface{} = 1

	i := a
	j, e := i.(int)
	if !e {
		fmt.Println("j=", j)
	}

	fmt.Println(i)
	fmt.Println(j)
}

func StudyInterface() {

	r := Rect1{10, 20}
	c := Circle1{10}

	showArea(r, c)

	Assertion()

	var test interface{} = 1
	switch test.(type) {
	case int:
		fmt.Println("int 네")
	}

}
