package utility

import (
	"encoding/json"
	"fmt"
)

type Member struct {
	Name   string
	Age    int
	Active bool
}

func studyJson_1() {

	mem := Member{"Alex", 10, true}

	//struct -> []byte
	jsonBytes, err := json.Marshal(mem)
	if err != nil {
		panic(err)
	}

	//[]byte -> string
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)
}

func studyJson_2() {
	jsonBytes, _ := json.Marshal(Member{"Tim", 1, true})

	mem := new(Member)
	err := json.Unmarshal(jsonBytes, mem)
	if err != nil {
		panic(err)
	}

	fmt.Println(mem.Name, mem.Age, mem.Active)
	fmt.Printf("%#v\n", mem)
}

func studyJson_3() {
	jsonBytes, _ := json.Marshal(Member{"Tim", 1, true})

	mem := new(map[string]interface{})
	err := json.Unmarshal(jsonBytes, mem)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", mem)
}

func StudyJson() {
	studyJson_3()
}
