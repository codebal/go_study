package basic

import (
	"fmt"
)

type Rect struct {
	width, height int
}

func (rect Rect) area() int {
	return rect.width * rect.height
}

func (rect *Rect) area2() int {
	rect.width++
	return rect.width * rect.height
}

func StudyMethod() {
	rect := Rect{10, 10}
	area := rect.area2()
	fmt.Printf("%v , %v\n", area, rect)
}
