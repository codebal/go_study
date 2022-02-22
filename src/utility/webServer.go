package utility

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func studyWebServer_1() {
	defer func() {
		fmt.Println("bye~")
		if r := recover(); r != nil {
			fmt.Println("error : ", r)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		sig := <-c
		fmt.Println("clean up : ", sig)
		os.Exit(1)
	}()

	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":8080", nil)
	fmt.Println("start")
}

type testHandler struct {
	http.Handler
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	str := "Your Request Path is " + req.URL.Path
	w.Write([]byte(str))
}

func studyWebServer_2() {
	http.Handle("/", new(testHandler))

	http.ListenAndServe(":5000", nil)
}

func StudyWebServer() {
	studyWebServer_1()
}
