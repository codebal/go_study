package utility

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Member1 struct {
	Name   string
	Age    int
	Active bool
}

type Member1s struct {
	Member []Member1
}

//encodeing
func studyXml_1() {
	mem := Member1{"Alex", 10, true}

	xmlBytes, err := xml.Marshal(mem)
	if err != nil {
		panic(err)
	}

	xmlString := string(xmlBytes)
	fmt.Println(xmlString)
}

//decording
func studyXml_2() {
	xmlBytes, _ := xml.Marshal(Member{"Alex", 10, true})

	var mem *Member
	err := xml.Unmarshal(xmlBytes, mem)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", mem)
}

func studyXml_3() {
	fp, err := os.Open("./assets/xml1.xml")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	data, err := ioutil.ReadAll(fp)
	if err != nil {
		panic(err)
	}

	members := new(Member1s)
	xmlError := xml.Unmarshal(data, members)
	if xmlError != nil {
		panic(xmlError)
	}
	fmt.Printf("%+v", members)
}

func StudyXml() {
	studyXml_3()
}
