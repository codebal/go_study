package utility

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func studyHttp_1() {
	defer func() {
		fmt.Println("defer first")

	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("defer recover")
		}
	}()

	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() //Response.Body.Close 는 항상 해줘야 하는구만

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(data))
}

func studyHttp_2() {
	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent", "Crawler")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Body)
	str := string(bytes)
	_ = str
	fmt.Println(str)
}

//post 로 일반 text 전송
func studyHttp_3() {
	reqBody := bytes.NewBufferString("post plain text")
	resp, err := http.Post("http://httpbin.org/post", "application/json", reqBody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		fmt.Println(str)
	}
}

//post로 form 데이터 전송
func studyHttp_4() {
	resp, err := http.PostForm("http://httpbin.org/post", url.Values{
		"name": {"Lee"},
		"Age":  {"10"},
	})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}

//json 사용
func studyHttp_5() {
	type Person struct {
		Name string
		Age  int
	}

	person := Person{"Kim", 46}
	pbytes, _ := json.Marshal(person)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("http://httpbin.org/post", "application/json", buff)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		fmt.Println(str)
	}
}

//xml 사용
func studyHttp_6() {
	type Person struct {
		Name string
		Age  int
	}

	person := Person{"haha", 40}
	pbytes, _ := xml.Marshal(person)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("http://httpbin.org/post", "application/json", buff)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(string(respBody))
	}
}

//Request 객체
func studyHttp_7() {
	type Person struct {
		Name string
		Age  int
	}

	person := Person{"yaho", 20}
	pbytes, _ := json.Marshal(person)
	buff := bytes.NewBuffer(pbytes)

	req, err := http.NewRequest("POST", "http://httpbin.org/post", buff)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(string(respBody))
	}
}

func StudyHttp() {
	studyHttp_3()
}
