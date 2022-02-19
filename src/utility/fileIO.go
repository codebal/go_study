package utility

import (
	"fmt"
	"io"
	"os"
)

func studyFileIO_1() {
	fi, err := os.Open("assets/1.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	fo, err := os.Create("assets/2.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	buff := make([]byte, 1024)

	for {
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if cnt == 0 {
			break
		}

		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("program end")
}

func StudyFileIO() {
	studyFileIO_1()
}
