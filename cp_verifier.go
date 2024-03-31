package main

import (
	"fmt"
)

func main() {
    fmt.Println(Hello(""))
}

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name
}