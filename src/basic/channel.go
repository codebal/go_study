package basic

import (
	"fmt"
	"time"
)

func studyChannel_1() {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		ch <- 123
	}()

	i := <-ch
	println(i)
}

func sendChan(ch chan<- string) {
	ch <- "Data"
}

func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}

func studyChannel_2() {
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 300)
			fmt.Println(i)
		}
		done <- true
	}()
	<-done
}

func studyChannel_3() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	close(ch)

	println(<-ch)
	println(<-ch)

	if _, ok := <-ch; !ok {
		println("없음")
	}
}

func studyChannel_4() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	close(ch)

	// for {
	// 	if i, ok := <-ch; ok {
	// 		println("i=", i)
	// 	} else {
	// 		break
	// 	}
	// }

	for i := range ch {
		println("i = ", i)
	}
}

func studyChannel_5() {
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1:
			println("run1 complete")
		case <-done2:
			println("run2 complete")
			break EXIT
		}
	}
}

func run1(done chan bool) {
	time.Sleep(time.Second * 3)
	println("before run1 done")
	done <- true
	println("after run1 done")
}

func run2(done chan bool) {
	time.Sleep(time.Second * 2)
	println("before run2 done")
	done <- true
	println("after run2 done")
}

func StudyChannel() {
	studyChannel_5()
}
