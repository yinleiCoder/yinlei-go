package main

import "fmt"

func main() {
	done := make(chan int, 1)
	go func() {
		fmt.Println("Hello, World!")
		done <- 1
	}()
	<-done
}
