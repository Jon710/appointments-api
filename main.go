package main

import (
	"appointments-api/mypackage"
	"fmt"
	"log"
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
	message, err := mypackage.PrintHello("João")

	if err != nil {
		log.Fatal(err)
	}

	p := Person{Name: "João", Age: 26}

	fmt.Println(message)
	p.PrintHelloWithStruct()

	// http.ListenAndServe(":8080", nil)
}
