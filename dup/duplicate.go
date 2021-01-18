package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Duplicate struct {
	filename string
	line     string
}

func main() {
	counts := make(map[Duplicate]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
		} else {
			for _, line := range strings.Split(string(data), "\n") {
				if line != "\n" {
					counts[Duplicate{filename, line}]++
				}
			}
		}
	}
	for duplicate, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s from %s\n", n, duplicate.line, duplicate.filename)
		}
	}
}
