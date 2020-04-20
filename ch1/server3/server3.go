package main

import (
	"fmt"
	l "github.com/davidgarciaMontreal/golang/ch1/lissajous"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/lisa", lisa)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lisa(w http.ResponseWriter, r *http.Request) {
	c := 5
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	if val, ok := r.Form["cycle"]; ok {
		s := strings.Join(val, "")
		c, _ = strconv.Atoi(s)
	}
	l.Lissajous(w, c)
}
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] %q\n", k, v)
	}

	fmt.Println("in handler")
}
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "URL.Path=%q count=%d\n", r.URL.Path, count)
	mu.Unlock()
	fmt.Println("in /count")
}
