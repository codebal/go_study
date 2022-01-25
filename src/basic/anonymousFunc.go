package main

func main() {
	{
		//익명함수
		sum := func(n ...int) int {
			s := 0
			for _, i := range n {
				s += i
			}
			return s
		}

		result := sum(1, 2, 3, 4, 5)
		println("result=", result)
	}

	{
		//일급함수
		add := func(i int, j int) int {
			return i + j
		}

		r1 := calc(add, 10, 20)
		println("r1 = ", r1)

		r2 := calc(func(x, y int) int { return x - y }, 10, 20)
		println("r2 = ", r2)
	}

	{
		//원형함수. type 정의
		calc := func(f calculator, a int, b int) int {
			result := f(a, b)
			return result
		}
		r := calc(func(a, b int) int { return a + b }, 10, 20)
		println("r == ", r)
	}
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

type calculator func(int, int) int
