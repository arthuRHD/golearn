package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func basic_echo() string {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo_inline() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}

func main() {
	echo_inline()
	start := time.Now()
	fmt.Println(basic_echo())
	end := time.Now()
	elapsed := end.Sub(start)
	println(elapsed)
	new_start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	new_end := time.Now()
	new_elapsed := new_end.Sub(new_start)
	println(new_elapsed)
}
