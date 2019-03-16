// Dup は標準入力から2回以上現れる業を出現回数と一緒に表示します。
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func getFileName(line string) []string {
	return strings.Split(line, "/")
}

func main() {
	counts := make(map[string]int)
	for _, fileName := range os.Args[1:] {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line+"/"+path.Base(fileName)]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			keyName := getFileName(line)
			fmt.Printf("[FileName: %s] %d\t%s\n", keyName[1], n, keyName[0])
		}
	}
}
