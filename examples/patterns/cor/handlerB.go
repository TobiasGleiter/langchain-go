package main

import (
	"fmt"
)

type ConcreteHandlerB struct {
	next Handler
}

func (h *ConcreteHandlerB) SetNext(handler Handler) {
	h.next = handler
}

func (h *ConcreteHandlerB) Handle(request string) error {
	if request == "B" {
		fmt.Println("ConcreteHandlerB handled request")
		return nil
	} else if h.next != nil {
		return h.next.Handle(request)
	} else {
		return fmt.Errorf("ConcreteHandlerB: unhandled request: %s", request)
	}
}
