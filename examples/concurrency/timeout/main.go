package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("Joe")
	timeout := time.After(5 * time.Second)

	for {
		select {
		case c := <-c:
			fmt.Println(c)
		case <-timeout:
			fmt.Println("You talk to much...")
			return
		}
	}
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}
