package basic

func studyIterate() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println("sum= ", sum)

	n := 1
	for n < 100 { //for ;n < 100; {
		n *= 2
		println("n=", n)
	}

	// for {
	// 	time.Sleep(time.Millisecond * 500)
	// 	fmt.Printf("Infinite loop, %v\n", time.Now())
	// }
	// 무한루프

	//range
	names := []string{"홍길동", "이순신", "강감찬"}
	for i, name := range names {
		println(i, name)
	}

	//goto
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue
		}
		a++
		if a > 10 {
			break
		}
	}
	if a == 11 {
		goto END
	}
	a += 100

END:
	println("a = ", a)
	println("End")

	//break label
	i := 0

L1:
	for {
		i++
		if i == 10 {
			break L1 //L1 바로 밑에있는 for 루프를 빠져 나간다는 뜻
		}
		println("i <= ", i)
	}
	println("OK == ", i)
}
