package main

import "fmt"

func main() {
	fmt.Println(Hello("Naidu"))
}

func Hello(str string) string{
	return "Hello " + str
}

