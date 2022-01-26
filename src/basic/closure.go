package main

import (
	"fmt"
	"testing"
)

func nextValue() (f1 func() int) {
	i := 0
	f1 = func() int {
		i++
		return i
	}
	return
}

func TestClosure(t *testing.T) {
	next := nextValue()

	for i := 0; i < 3; i++ {
		println("next = ", next())
	}

	f1 := func() {
		println("f1")
	}

	fmt.Printf("f1 %T\n", f1)

}
