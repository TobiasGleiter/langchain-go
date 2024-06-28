package main

import (
	"fmt"
)

type ConcreteHandlerA struct {
	next Handler
}

func (h *ConcreteHandlerA) SetNext(handler Handler) {
	h.next = handler
}

func (h *ConcreteHandlerA) Handle(request string) error {
	if request == "A" {
		fmt.Println("ConcreteHandlerA handled request")
		return nil
	} else if h.next != nil {
		return h.next.Handle(request)
	} else {
		return fmt.Errorf("ConcreteHandlerA: unhandled request: %s", request)
	}
}
