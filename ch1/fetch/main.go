package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func addScheme(url string) (result string) {
	if !strings.HasPrefix(url, "http") {
		result = "http://" + url
	}
	return
}
func main() {
	for _, url := range os.Args[1:] {
		url = addScheme(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("StatusCode = %d\n", resp.StatusCode)
		fmt.Printf("%s", b)
	}
}
