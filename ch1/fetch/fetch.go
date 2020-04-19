package main

import (
	"fmt"
	"io"
	// "io/ioutil"
	"bytes"
	"net/http"
	"os"
	"strings"
)

const httpPrefix string = "http://"

func main() {
	for _, url := range os.Args[1:] {
		var theUrl bytes.Buffer
		// func HasPrefix(s, prefix string) bool
		if strings.HasPrefix(url, httpPrefix) {
			theUrl.WriteString(url)
		} else {
			theUrl.WriteString(httpPrefix)
			theUrl.WriteString(url)
		}
		resp, err := http.Get(theUrl.String())

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		n, err := io.Copy(os.Stdout, resp.Body)

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}
		fmt.Printf("\nNumber of bytes %d\n", n)
		fmt.Printf("Status Code %d\n", resp.StatusCode)
	}
}
