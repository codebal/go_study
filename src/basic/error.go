package basic

import (
	"fmt"
	"log"
	"os"
)

type MyError interface {
	Error() string
	TestFunc()
}

type MyErrorStruct struct {
	Name string
}

func (s *MyErrorStruct) Error() string {
	return "error = " + s.Name
}

func (s *MyErrorStruct) TestFunc() {
	fmt.Println("TestFunc")
}

func otherFunc() (string, error) {
	return "", &MyErrorStruct{"에라"}
}

func StudyError() {
	{
		wd, _ := os.Getwd()
		fmt.Println("wd = ", wd)
		f, err := os.Open("./basic/nil.go")
		if err != nil {
			log.Fatal(err.Error())
		}
		println(f.Name())
	}
	{
		_, err := otherFunc()
		switch err.(type) {
		default:
			fmt.Println("no error")
		case MyError:
			fmt.Println("error", err)
		case error:
			fmt.Println("error")
		}
	}
}
