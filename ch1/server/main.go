package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler はリクエストされたURL r のPath要素を返します。
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler) // ここのリクエストに対して handler が呼ばれる
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
