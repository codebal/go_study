package basic

import (
	"fmt"
	"os"
)

func StudyDeferPanic() {
	defer func() {
		fmt.Println("여긴 마지막에 무조껀 실행")
	}()

	f, err := os.Open("./basic/nil.go")
	if err != nil {
		panic(err)
	}

	defer func() {
		fmt.Println("이거 실행")
		f.Close()
	}()

	bytes := make([]byte, 1024)
	f.Read(bytes)
	println(len(bytes))
}

func StudyDeferPanicRecover() {
	defer fmt.Println("defer 1")
	openFile()
	defer fmt.Println("defer 2")

	println("Done")
}

func openFile() {
	defer func() { //recover는 현제 블록에서만
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR = ", r)
		}
	}()

	f, err := os.Open("./basic/nil.go1")
	if err != nil {
		panic(err)
	}

	defer func() {
		println("파일 닫기")
		f.Close()
	}()
}
