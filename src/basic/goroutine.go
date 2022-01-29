package basic

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func say1(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func studyGoroutine_1() {
	say1("Sync")

	go say1("Async1")
	go say1("Async2")
	go say1("Async3")

	time.Sleep(time.Second * 3)
}

func studyGoroutine_2() {
	var wait sync.WaitGroup
	wait.Add(2)

	go func() {
		rInt := time.Second * time.Duration(rand.Intn(5))
		fmt.Println("rInt=", rInt)
		time.Sleep(rInt)
		//time.Sleep(0)
		fmt.Println("Hello")
		wait.Done()
	}()

	go func(msg string) {
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		fmt.Println(msg)
		wait.Done()
	}("Hi")

	wait.Wait()
}

func studyGoroutine_3() {
	runtime.GOMAXPROCS(4)

	fmt.Println("NumCPU()=", runtime.NumCPU())

	var wait sync.WaitGroup
	wait.Add(10)

	for i := 0; i < 10; i++ {
		go func(name string) {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			fmt.Println(name)
			wait.Done()
		}(fmt.Sprintf("go-%v", i))
	}

	wait.Wait()
}

//go routine은 생성된 블록은 완료되도 끝까지 처리된다
//클로저로 인해서 블록의 참조된 값도 계속 유지된다.
func studyGoroutine_4() {
	func() {
		cnt := 0
		run1 := func(name string, secs int) {
			time.Sleep(time.Second * time.Duration(secs))
			cnt++
			fmt.Printf("name=%v, cnt=%v\n", name, cnt)
		}

		go run1("go1", 1)
		go run1("go2", 3)
		go run1("go3", 6)
		go run1("go4", 9)

		time.Sleep(time.Second * 5)
	}()

	//time.Sleep(time.Second * 5)

	fmt.Scanln()
}

func StudyGoroutine() {
	//studyGoroutine_1()
	studyGoroutine_4()
}
