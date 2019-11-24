package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("Naidu"))
}

/// "Hello String Function"
func Hello(str string) string {
	if str == "" {
		str = "World"
	}
	return englishHelloPrefix + str
}
