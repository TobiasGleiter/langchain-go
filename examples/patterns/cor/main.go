package main

import "fmt"

func main() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	handlerA.SetNext(handlerB)

	if err := handlerA.Handle("A"); err != nil {
		fmt.Println("Error:", err)
	}

	if err := handlerA.Handle("B"); err != nil {
		fmt.Println("Error:", err)
	}
}
