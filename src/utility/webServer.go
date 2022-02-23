package utility

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
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
		fmt.Printf("ready channel %#v\n", c)
		sig := <-c
		fmt.Printf("clean up : %#v, %#v\n", sig, c)
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

	http.ListenAndServe(":8080", nil)
}

type staticHandler struct {
}

func (h *staticHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	localPath := "wwwroot" + req.URL.Path
	fmt.Printf("%v\n", localPath)
	content, err := ioutil.ReadFile(localPath)
	fmt.Println("content", string(content))
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		fmt.Println(http.StatusText(404))
		return
	}
	contentType := getContentType(localPath)
	fmt.Println("contentType", contentType)
	w.Header().Add("Content-Type", contentType)
	w.Write(content)
}

func getContentType(localPath string) (contentType string) {
	ext := filepath.Ext(localPath)

	switch ext {
	case ".html":
		contentType = "text/html; charset=utf-8"
	case ".css":
		contentType = "text/css; charset=utf-8"
	case ".js":
		contentType = "application/javascript; charset=utf-8"
	case ".png":
		contentType = "image/png"
	case ".jpg":
		contentType = "image/jpeg"
	default:
		contentType = "text/plain; charset=utf-8"
	}

	return contentType
}

func studyWebServer_3() {
	//static file server custom
	http.Handle("/", new(staticHandler))
	http.ListenAndServe(":8080", nil)
}

func studyWebServer_4() {
	http.Handle("/", http.FileServer(http.Dir("wwwroot")))
	fmt.Println(http.Dir("wwwroot"))
	http.ListenAndServe(":8080", nil)
}

func StudyWebServer() {
	studyWebServer_4()
}
