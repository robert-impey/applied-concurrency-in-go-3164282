package main

import (
	"fmt"
	"time"
)

func main() {
	go hello()
	time.Sleep(time.Second * 1)
	goodbye()
}

func hello() {
	fmt.Println("Hello, world!")
}

func goodbye() {
	fmt.Println("Goodbye, world!")
}
