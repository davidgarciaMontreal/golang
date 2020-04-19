package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	name := RandStringRunes(10)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	// go doc os.File | grep Write
	// 	func (f *File) Write(b []byte) (n int, err error)
	// f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	// go doc http.Response.Write -> func (r *Response) Write(w io.Writer) error

	// go doc io.Writer
	//  type Writer interface {
	//	Write(p []byte) (n int, err error)
	//    }
	// nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	fileName := fmt.Sprintf("file-%s.txt", name)

	f, errr := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if errr != nil {
		ch <- fmt.Sprintf("while opening file.txt %v", errr)
		return
	}
	defer f.Close()
	err = resp.Write(f)
	if err != nil {
		ch <- fmt.Sprintf("while writing %s: %v", url, err)
		return
	}
	fi, err := f.Stat()
	if err != nil {
		ch <- fmt.Sprintf("f.Stat failed %s: %v", url, err)
		return
	}
	nbytes := fi.Size()
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s %s", secs, nbytes, url, fileName)
}
