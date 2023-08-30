package main

import (
	"fmt"
	// "net/http"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) PrintHelloWithStruct() {
	fmt.Printf("Hello, my name is %s and I am %d years old\n", p.Name, p.Age)
}

func main() {
	p := Person{Name: "João", Age: 26}

	p.PrintHelloWithStruct()

	// http.ListenAndServe(":8080", nil)
}
