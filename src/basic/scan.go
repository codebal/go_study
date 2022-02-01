package basic

import "fmt"

func studyScan_1() {
	var (
		name, addr string
		num        int
	)
	fmt.Print("name : ")
	fmt.Scanln(&name)
	fmt.Print("번호 : ")
	fmt.Scanln(&num)
	fmt.Print("주소 : ")
	fmt.Scanln(&addr)
	fmt.Println("name : ", name, " num : ", num, " address : ", addr)
}

func studyScan_2() {
	var (
		name, addr string
		num        int
	)
	fmt.Print("name and num and address : ")
	re, err := fmt.Scanln(&name, &num, &addr)
	fmt.Println("name : ", name, " num : ", num, " address : ", addr)
	fmt.Println("scan result : ", re, " error : ", err)
}

//Scan 함수는 코드상의 모든 scan 값이 채워질때 까지 계속 값을 받고 종료되지 않는다.
func studyScan_3() {
	var (
		name, addr string
		num        int
	)
	fmt.Print("name and num and address : ")
	re, err := fmt.Scan(&name, &num, &addr)
	fmt.Println("name : ", name, " num : ", num, " address : ", addr)
	fmt.Println("scan result : ", re, " error : ", err)
}

func studyScan_4() {
	var n1, n2, n3, n4 int
	fmt.Println("IPv4 주소")
	fmt.Scanf("%d.%d.%d.%d", &n1, &n2, &n3, &n4)
	fmt.Printf("input value is : %d.%d.%d.%d \n", n1, n2, n3, n4)
}

func StudyScan() {
	studyScan_4()
}
