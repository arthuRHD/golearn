package main

import (
	"fmt"
	"log"
)

func display(msg string) string {
	return msg + " from another package"
}

func main() {
	log.Println(Hello("Arthur"))
	log.Println(display(Hello("Arthur")))
}

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
