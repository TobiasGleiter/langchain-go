package main

type Handler interface {
	SetNext(handler Handler)
	Handle(request string) error
}
